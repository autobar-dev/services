package main

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	echo "github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"

	"github.com/autobar-dev/services/user/middleware"
	"github.com/autobar-dev/services/user/repositories"
	"github.com/autobar-dev/services/user/routes"
	"github.com/autobar-dev/services/user/types"
	"github.com/autobar-dev/services/user/utils"
	authrepository "github.com/autobar-dev/shared-libraries/go/auth-repository"
	emailrepository "github.com/autobar-dev/shared-libraries/go/email-repository"
	emailtemplaterepository "github.com/autobar-dev/shared-libraries/go/emailtemplate-repository"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("failed to load .env file", err)
	}

	config, err := utils.LoadEnvVars()
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

	e := echo.New()
	e.HideBanner = true

	app_context := &types.AppContext{
		MetaFactors: utils.GetMetaFactors(),
		Config:      config,
		Repositories: &types.Repositories{
			Cache:                  repositories.NewCacheRepository(redis_client),
			User:                   repositories.NewUserRepository(database),
			UnfinishedRegistration: repositories.NewUnfinishedRegistrationRepository(database),
			Auth:                   authrepository.NewAuthRepository(config.AuthServiceURL, types.MicroserviceName),
			Email:                  emailrepository.NewEmailRepository(config.EmailServiceURL, types.MicroserviceName),
			EmailTemplate: emailtemplaterepository.NewEmailTemplateRepository(
				config.EmailTemplateServiceURL,
				types.MicroserviceName,
			),
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
	e.GET("/", routes.GetUserRoute)
	e.GET("/who-am-i", routes.WhoAmIRoute)
	e.POST("/create", routes.CreateUserRoute)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", (*config).Port)))
}
