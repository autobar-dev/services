package main

import (
	"fmt"

	"github.com/autobar-dev/services/email/providers"
	"github.com/autobar-dev/services/email/routes"
	"github.com/autobar-dev/services/email/types"
	"github.com/autobar-dev/services/email/utils"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("failed to load .env file: %s", err)
	}

	// Load environment variables
	config, err := utils.LoadEnvVars()
	if err != nil {
		panic(fmt.Sprintf("failed to load environment variables: %s", err))
	}

	// Set up logging
	logger := utils.GetLogger(config.LoggerEnvironment)
	defer logger.Sync()

	// Create app context
	app_context := &types.AppContext{
		Logger:       logger,
		MetaFactors:  utils.GetMetaFactors(),
		Config:       config,
		Repositories: &types.Repositories{},
		Providers: &types.Providers{
			Email: providers.NewSmtpEmailProvider(
				config.SmtpHostname,
				config.SmtpPort,
				config.SmtpUsername,
				config.SmtpPassword,
			),
		},
	}

	// Initialize HTTP server
	e := echo.New()
	e.HideBanner = true

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
	e.POST("/send", routes.SendRoute)

	logger.Fatal(e.Start(fmt.Sprintf(":%d", (*config).Port)))
}
