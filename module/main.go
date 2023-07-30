package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	echo "github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"

	"go.a5r.dev/services/module/repositories"
	"go.a5r.dev/services/module/routes"
	"go.a5r.dev/services/module/types"
	"go.a5r.dev/services/module/utils"
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

	database, err := sqlx.Connect("postgres", config.DatabaseURL)
	if err != nil {
		fmt.Println("failed to connect to database", err)
		os.Exit(1)
	}
	defer database.Close()

	redis_options, err := redis.ParseURL(config.RedisURL)
	redis_client := redis.NewClient(redis_options)
	defer redis_client.Close()

	amqp_connection, err := amqp091.Dial(config.AmqpURL)
	if err != nil {
		fmt.Println("failed to connect to queue", err)
		os.Exit(1)
	}
	defer amqp_connection.Close()

	amqp_channel, err := amqp_connection.Channel()
	if err != nil {
		fmt.Println("failed to open a channel in queue", err)
		os.Exit(1)
	}
	defer amqp_channel.Close()

	rand.Seed(time.Now().UnixNano())

	e := echo.New()
	e.HideBanner = true

	app_context := &types.AppContext{
		Meta:        utils.LoadMeta(),
		AmqpChannel: amqp_channel,
		Config:      config,
		Repositories: &types.Repositories{
			Module:   repositories.NewModuleRepository(database),
			Cache:    repositories.NewCacheRepository(redis_client),
			Auth:     repositories.NewAuthRepository(config.AuthServiceURL),
			Realtime: repositories.NewRealtimeRepository(config.RealtimeServiceURL),
		},
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
	e.GET("/", routes.GetModuleRoute)
	e.GET("/all", routes.GetAllModulesRoute)
	e.GET("/all-for-station", routes.GetAllModulesForStationRoute)
	e.GET("/request-report", routes.RequestReportRoute)
	e.POST("/create", routes.CreateModuleRoute)
	e.POST("/report", routes.ReportRoute)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", (*config).Port)))
}
