package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	app_grpc "github.com/autobar-dev/services/currency/grpc"
	"github.com/autobar-dev/services/currency/grpc/generated_grpc"
	"github.com/autobar-dev/services/currency/routes"
	"github.com/autobar-dev/services/currency/stores"
	"github.com/autobar-dev/services/currency/stores/postgres"
	"github.com/autobar-dev/services/currency/types"
	"github.com/autobar-dev/services/currency/types/interfaces"
	"github.com/charmbracelet/log"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	var l interfaces.AppLogger = log.Default()

	// Load env file
	err := godotenv.Load(".env")

	if err == nil {
		l.Info("Loaded .env file")
	}

	cs := os.Getenv("DB_CONNECTION_STRING")
	db, err := sqlx.Connect("postgres", cs)

	if err != nil {
		l.Error("Error connecting to the database.")
		os.Exit(1)
	}

	// Initialize rate store
	rs, err := postgres.NewPostgresRateStore(&l, db)

	if err != nil {
		l.Error(err)
		os.Exit(1)
	}

	l.Info("Connected to the database.")

	// Initialize supported currencies store
	scs, err := postgres.NewPostgresSupportedCurrenciesStore(&l, db)

	if err != nil {
		l.Error(err)
		os.Exit(1)
	}

	// Initialize remote exchange rate store
	rers, err := stores.NewExchangeRateApiStore()

	if err != nil {
		l.Error(err)
		os.Exit(1)
	}

	// Initialize stores
	stores := types.AppStores{
		RateStore:                rs,
		SupportedCurrenciesStore: scs,

		RemoteExchangeRateStore: rers,
	}

	// Create app context
	app_context := types.AppContext{
		AppLogger: &l,
		Stores:    &stores,
	}

	// Initialize REST
	e := echo.New()

	rp := os.Getenv("REST_PORT")
	es := http.Server{
		Addr:    ":" + rp,
		Handler: e,
	}

	// Mount context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			rc := types.RestContext{
				Context:    c,
				AppContext: &app_context,
			}

			return next(rc)
		}
	})

	// Attach route handlers
	e.GET("/", routes.Currency)
	e.GET("/supported", routes.Supported)
	e.GET("/rate", routes.Rate)

	e.POST("/create", routes.Create)

	e.PUT("/set-enabled", routes.SetEnabled)
	e.PUT("/set-rate", routes.SetRate)
	e.PUT("/force-update-rate", routes.ForceUpdateRate)

	e.DELETE("/delete", routes.Delete)

	// REST: start listening
	go func() {
		l.Info(fmt.Sprintf("REST listening on port %s...", rp))

		if err := es.ListenAndServe(); err != http.ErrServerClosed {
			l.Error("REST: Unable to listen.", "error", err)
			os.Exit(1)
		}
	}()

	// Set up GRPC
	gs := grpc.NewServer()
	ch := app_grpc.NewCurrencyHandler(&app_context)

	generated_grpc.RegisterCurrencyServer(gs, ch)
	reflection.Register(gs)

	gp := os.Getenv("GRPC_PORT")
	gl, err := net.Listen("tcp", ":"+gp)

	if err != nil {
		l.Error("GRPC: unable to listen", "error", err)
		os.Exit(1)
	}

	// GRPC: start
	go func() {
		l.Info(fmt.Sprintf("GRPC listening on port %s...", gp))

		gs.Serve(gl)
	}()

	// Handling shutdown
	signal_channel := make(chan os.Signal, 1)
	signal.Notify(signal_channel, syscall.SIGINT, syscall.SIGTERM)

	<-signal_channel

	l.Info("Received termination signal.")

	// Shut down REST server
	l.Info("Shutting down REST server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		l.Error("Could not shut down REST server gracefully.")
		os.Exit(1)
	}

	l.Info("Gracefully shut down REST server.")

	// Shut down GRPC server
	l.Info("Shutting down GRPC server...")
	gs.GracefulStop()

	l.Info("Gracefully shut down GRPC server.")
}
