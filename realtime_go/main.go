package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	echo "github.com/labstack/echo/v4"
	sse "github.com/r3labs/sse/v2"
	amqp "github.com/rabbitmq/amqp091-go"

	"go.a5r.dev/services/realtime/repositories"
	"go.a5r.dev/services/realtime/routes"
	"go.a5r.dev/services/realtime/types"
	"go.a5r.dev/services/realtime/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("failed to load .env file", err)
	}

	config, err := types.LoadEnvVars()
	if err != nil {
		fmt.Println("failed to load .env vars", err)
		os.Exit(1)
	}

	amqp_connection, err := amqp.Dial(config.AmqpURL)
	if err != nil {
		fmt.Println("failed to connect to amqp", err)
		os.Exit(1)
	}
	defer amqp_connection.Close()

	amqp_channel, err := amqp_connection.Channel()
	if err != nil {
		fmt.Println("failed to open a channel in queue", err)
		os.Exit(1)
	}
	defer amqp_channel.Close()

	sse_server := sse.New()
	defer sse_server.Close()

	e := echo.New()
	e.HideBanner = true

	app_context := &types.AppContext{
		Meta:        utils.LoadMeta(),
		AmqpChannel: amqp_channel,
		Config:      config,
		Repositories: &types.Repositories{
			Auth: repositories.NewAuthRepository(config.AuthServiceURL),
		},
		SseServer: sse_server,
	}

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			rest_context := &types.RestContext{
				c,
				app_context,
			}
			return next(rest_context)
		}
	})

	e.GET("/meta", routes.MetaRoute)
	e.GET("/events", routes.EventsRoute)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", (*config).Port)))
}
