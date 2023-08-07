package main

import (
	"fmt"

	"github.com/autobar-dev/services/auth/routes"
	oauth_routes "github.com/autobar-dev/services/auth/routes/oauth"
	"github.com/autobar-dev/services/auth/types"
	"github.com/autobar-dev/services/auth/utils"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
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

	// Connect to database
	database, err := sqlx.Connect("postgres", config.DatabaseURL)
	if err != nil {
		logger.Fatalf("failed to connect to database", err)
	}
	defer database.Close()

	// Create app context
	app_context := &types.AppContext{
		Logger:       logger,
		MetaFactors:  utils.GetMetaFactors(),
		Config:       config,
		Repositories: &types.Repositories{},
	}

	// Initialize OAuth
	// oauth_manager, err := utils.SetupOAuthManager(app_context)
	// if err != nil {
	// 	logger.Fatalf("failed to set up OAuth manager", err)
	// }
	//
	// oauth_server, err := utils.SetupOAuthServer(app_context, oauth_manager)
	// if err != nil {
	// 	logger.Fatalf("failed to set up OAuth server", err)
	// }
	// app_context.Repositories.OAuthServer = oauth_server

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

	e.Any("/oauth/authorize", oauth_routes.AuthorizeRoute)
	e.POST("/oauth/token", oauth_routes.TokenRoute)

	logger.Fatal(e.Start(fmt.Sprintf(":%d", (*config).Port)))
}
