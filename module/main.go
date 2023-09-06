package main

import (
	"fmt"
	"os"

	authrepository "github.com/autobar-dev/shared-libraries/go/auth-repository"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"

	"github.com/autobar-dev/services/module/middleware"
	"github.com/autobar-dev/services/module/repositories"
	"github.com/autobar-dev/services/module/routes"
	"github.com/autobar-dev/services/module/types"
	"github.com/autobar-dev/services/module/utils"
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

	redis_cache_options, err := redis.ParseURL(config.RedisCacheURL)
	redis_cache_client := redis.NewClient(redis_cache_options)
	defer redis_cache_client.Close()

	redis_state_options, err := redis.ParseURL(config.RedisStateURL)
	redis_state_client := redis.NewClient(redis_state_options)
	defer redis_state_client.Close()

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

	e := echo.New()
	e.HideBanner = true

	app_context := &types.AppContext{
		MetaFactors: utils.GetMetaFactors(),
		AmqpChannel: amqp_channel,
		Config:      config,
		Repositories: &types.Repositories{
			Module:   repositories.NewModuleRepository(database),
			Cache:    repositories.NewCacheRepository(redis_cache_client),
			State:    repositories.NewStateRepository(redis_state_client),
			Realtime: repositories.NewRealtimeRepository(config.RealtimeServiceURL),
			Auth:     authrepository.NewAuthRepository(config.AuthServiceURL, types.MicroserviceName),
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
	e.GET("/", routes.GetModuleRoute)
	e.GET("/all", routes.GetAllModulesRoute)
	e.GET("/all-for-station", routes.GetAllModulesForStationRoute)
	e.GET("/request-report", routes.RequestReportRoute)
	e.POST("/create", routes.CreateModuleRoute)
	e.POST("/report", routes.ReportRoute)
	e.POST("/activate", routes.ActivateRoute)
	e.POST("/deactivate", routes.DeactivateRoute)
	e.PATCH("/update-activation-session", routes.UpdateActivationSessionRoute)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", (*config).Port)))
}
