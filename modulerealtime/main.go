package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/autobar-dev/services/modulerealtime/handlers"
	"github.com/autobar-dev/services/modulerealtime/types"
	"github.com/autobar-dev/services/modulerealtime/types/interfaces"
	"github.com/autobar-dev/services/modulerealtime/utils"
	"github.com/charmbracelet/log"
	socketio "github.com/googollee/go-socket.io"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	// Initialize logger
	charm_logger := log.Default()
	charm_logger.SetLevel(log.DebugLevel)

	var logger interfaces.AppLogger = charm_logger

	// Load env file
	err := godotenv.Load(".env")

	if err == nil {
		logger.Debug("Loaded .env file")
	}

	// Check env variables
	missing_env_vars := utils.CheckEnvs()

	if len(missing_env_vars) > 0 {
		logger.Error("Missing environment variables", "missing_env_vars", missing_env_vars)
		os.Exit(1)
	}

	// Create context
	app_context := types.AppContext{
		AppLogger: &logger,
	}

	// Initialize socket.io server
	socketio_server := socketio.NewServer(nil)

	// Initialize redis adapter
	redis_addr := os.Getenv("REDIS_ADDRESS")
	redis_password := os.Getenv("REDIS_PASSWORD")

	_, err = socketio_server.Adapter(&socketio.RedisAdapterOptions{
		Addr:     redis_addr,
		Password: redis_password,
	})

	if err != nil {
		logger.Error("error connecting to Redis", "err", err)
		os.Exit(1)
	}

	logger.Debug("Connected to Redis")

	// Register socket.io events
	socketio_server.OnConnect("/", func(socket socketio.Conn) error {
		return handlers.OnConnect(&app_context, socket)
	})

	socketio_server.OnEvent("/", "auth", func(s socketio.Conn, msg string) {
		s.SetContext(msg)
		s.Emit("yooo", strings.ToUpper(msg))
		logger.Info("root", "id", s.ID(), "msg", msg)
	})

	socketio_server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("yooo", last)
		s.Close()
		return last
	})

	socketio_server.OnError("/", func(s socketio.Conn, e error) {
		logger.Error("meet error:", "e", e)
	})

	socketio_server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		logger.Info("closed", "reason", reason)
	})
	go func() {
		if err := socketio_server.Serve(); err != nil {
			logger.Error("error serving socket.io", "err", err)
		}
	}()
	defer socketio_server.Close()

	// Initialize REST
	e := echo.New()
	e.HideBanner = true

	rest_port := os.Getenv("REST_PORT")
	rest_server := http.Server{
		Addr:    ":" + rest_port,
		Handler: e,
	}

	// Mount REST context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			rest_context := types.RestContext{
				Context:    c,
				AppContext: &app_context,
			}

			return next(rest_context)
		}
	})

	// Attach route handlers
	e.Any("/socket.io/", func(c echo.Context) error {
		socketio_server.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	// Start REST server
	go func() {
		logger.Debug(fmt.Sprintf("REST listening on port %s...", rest_port))

		if err := rest_server.ListenAndServe(); err != http.ErrServerClosed {
			logger.Error("REST: Unable to listen.", "error", err)
			os.Exit(1)
		}
	}()

	// Handling shutdown
	signal_channel := make(chan os.Signal, 1)
	signal.Notify(signal_channel, syscall.SIGINT, syscall.SIGTERM)

	<-signal_channel

	logger.Info("Received termination signal.")

	// Shut down REST server
	logger.Info("Shutting down REST server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		logger.Error("REST: Unable to shut down.", "error", err)
		os.Exit(1)
	}

	logger.Info("Gracefully shut down REST server.")
}
