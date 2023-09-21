package main

import (
	"fmt"

	"github.com/autobar-dev/shared-libraries/go/auth-repository"
	"github.com/joho/godotenv"
	echo "github.com/labstack/echo/v4"
	sse "github.com/r3labs/sse/v2"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"

	"github.com/autobar-dev/services/realtime/middleware"
	"github.com/autobar-dev/services/realtime/repositories"
	"github.com/autobar-dev/services/realtime/routes"
	"github.com/autobar-dev/services/realtime/types"
	"github.com/autobar-dev/services/realtime/utils"
)

func main() {
	_ = godotenv.Load()

	config, err := types.LoadEnvVars()
	if err != nil {
		panic(err)
	}

	amqp_connection, err := amqp.Dial(config.AmqpURL)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to amqp: %s", err))
	}
	defer amqp_connection.Close()

	amqp_channel, err := amqp_connection.Channel()
	if err != nil {
		panic(fmt.Sprintf("failed to open a channel in queue: %s", err))
	}
	defer amqp_channel.Close()

	redis_state_options, err := redis.ParseURL(config.RedisStateURL)
	if err != nil {
		panic(fmt.Sprintf("failed to open redis: %s", err))
	}

	redis_state_client := redis.NewClient(redis_state_options)

	sse_server := sse.New()
	defer sse_server.Close()

	e := echo.New()
	e.HideBanner = true

	app_context := &types.AppContext{
		MetaFactors: utils.GetMetaFactors(),
		AmqpChannel: amqp_channel,
		Config:      config,
		Repositories: &types.Repositories{
			Auth:  authrepository.NewAuthRepository(config.AuthServiceURL, types.MicroserviceName),
			State: repositories.NewStateRepository(redis_state_client),
			Mq:    repositories.NewMqRepository(amqp_channel),
		},
	}

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			rest_context := &types.RestContext{
				c,
				app_context,
				nil,
			}
			return next(rest_context)
		}
	})

	e.Use(middleware.AccessTokenMiddleware)

	e.GET("/meta", routes.MetaRoute)
	e.GET("/ws", routes.WsRoute)
	e.POST("/send-command", routes.SendCommandRoute)
	e.POST("/reply", routes.ReplyRoute)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", (*config).Port)))
}
