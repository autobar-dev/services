package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/autobar-dev/services/modulerealtime/types"
	"github.com/autobar-dev/services/modulerealtime/utils"
	"github.com/autobar-dev/services/modulerealtime/views/handlers"
	"github.com/autobar-dev/services/modulerealtime/views/routes"

	"github.com/charmbracelet/log"
	socketio "github.com/googollee/go-socket.io"
	"github.com/joho/godotenv"
	echo "github.com/labstack/echo/v4"
)

func main() {
	hash := "no-commit"
	version := "1.2.3"

	log := log.Default()

	err := godotenv.Load(".env")

	if err != nil {
		log.Debug("Loaded environment variables from file")
	}

	config, err := utils.LoadConfig()

	if err != nil {
		log.Fatal("Error loading config from env variables", "error", err)
	}

	app_ctx := types.AppContext{
		Log: log,
		Meta: &types.Meta{
			Hash:    hash,
			Version: version,
		},
	}

	socket_server := socketio.NewServer(nil)

	socket_server.OnConnect("/", func(s socketio.Conn) error {
		return handlers.OnConnect(s, &app_ctx)
	})
	socket_server.OnDisconnect("/", handlers.OnDisconnect)

	go socket_server.Serve()
	defer socket_server.Close()

	e := echo.New()
	e.HideBanner = true

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			rest_ctx := types.RestContext{
				Context:    c,
				AppContext: &app_ctx,
			}

			return next(&rest_ctx)
		}
	})

	e.Any("/socket.io/", func(context echo.Context) error {
		socket_server.ServeHTTP(context.Response(), context.Request())
		return nil
	})

	e.GET("/meta", routes.Meta)

	es := http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: e,
	}

	go func() {
		log.Info("HTTP server listening...", "port", config.Port)

		if err := es.ListenAndServe(); err != http.ErrServerClosed {
			log.Error("Unable to listen", "error", err)
			os.Exit(1)
		}
	}()

	signal_channel := make(chan os.Signal, 1)
	signal.Notify(signal_channel, syscall.SIGINT, syscall.SIGTERM)

	<-signal_channel

	log.Info("Received termination signal")

	// Shut down REST server
	log.Info("Shutting down HTTP server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Error("Could not shut down HTTP server gracefully")
		os.Exit(1)
	}

	log.Info("Gracefully shut down HTTP server")

}
