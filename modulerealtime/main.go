package main

import (
	"net/http"
	"os"
	"os/signal"

	"github.com/charmbracelet/log"
	socketio "github.com/googollee/go-socket.io"
	"github.com/joho/godotenv"
	echo "github.com/labstack/echo/v4"
)

func main() {
	log := log.Default()

	err := godotenv.Load(".env")

	if err != nil {
		log.Debug("Loaded environment variables from file")
	}

	port := os.Getenv("PORT")

	socket_server := socketio.NewServer(nil)

	go socket_server.Serve()
	defer socket_server.Close()

	e := echo.New()
	e.HideBanner = true

	e.Any("/socket.io/", func(context echo.Context) error {
		socket_server.ServeHTTP(context.Response(), context.Request())
		return nil
	})

	es := http.Server{
		Addr:    ":" + port,
		Handler: e,
	}

	go func() {
		log.Info("HTTP server listening...", "port", port)

		if err := es.ListenAndServe(); err != http.ErrServerClosed {
			log.Error("Unable to listen", "error", err)
			os.Exit(1)
		}
	}()

	signal_channel := make(chan os.Signal, 1)
	signal.Notify(signal_channel, syscall.SIGINT, syscall.SIGTERM)

	<-signal_channel

	log.Info("Received termination signal.")

	// Shut down REST server
	log.Info("Shutting down REST server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Error("Could not shut down HTTP server gracefully.")
		os.Exit(1)
	}

	l.Info("Gracefully shut down REST server.")

}
