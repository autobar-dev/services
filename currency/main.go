package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/autobar-dev/services/currency/routes"
	"github.com/autobar-dev/services/currency/stores/postgres"
	"github.com/autobar-dev/services/currency/types"
	"github.com/autobar-dev/services/currency/types/interfaces"
	"github.com/charmbracelet/log"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

func main() {
	var l interfaces.AppLogger = log.Default()

	// Load env file
	err := godotenv.Load(".env")

	if err != nil {
		l.Error("Not able to load .env file")
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

	// Initialize stores
	stores := types.AppStores{
		RateStore:                rs,
		SupportedCurrenciesStore: scs,
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
				Context:   c,
				AppLogger: &l,
				Stores:    &stores,
			}

			return next(rc)
		}
	})

	// Attach route handlers
	e.GET("/supported", routes.Supported)

	// REST: start listening
	go func() {
		l.Info(fmt.Sprintf("REST: Starting to listen on port %s...", rp))

		if err := es.ListenAndServe(); err != http.ErrServerClosed {
			l.Error("REST: Unable to listen.", "error", err)
			os.Exit(1)
		}
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
}
