package main

import (
	"fmt"

	"github.com/autobar-dev/services/auth/middleware"
	"github.com/autobar-dev/services/auth/providers"
	"github.com/autobar-dev/services/auth/repositories"
	"github.com/autobar-dev/services/auth/routes"
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

	// Create repositories
	auth_user_repository := repositories.NewUserRepository(database)
	auth_module_repository := repositories.NewModuleRepository(database)
	refresh_token_repository := repositories.NewRefreshTokenRepository(database)

	// Create auth provider
	postgres_auth_provider := providers.NewPostgresAuthProvider(
		auth_user_repository,
		auth_module_repository,
		refresh_token_repository,
		config.JwtSecret,
	)

	// Create app context
	app_context := &types.AppContext{
		Logger:      logger,
		MetaFactors: utils.GetMetaFactors(),
		Config:      config,
		Repositories: &types.Repositories{
			AuthUser:     auth_user_repository,
			AuthModule:   auth_module_repository,
			RefreshToken: refresh_token_repository,
		},
		Providers: &types.Providers{
			Auth: postgres_auth_provider,
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
				nil,
			}
			return next(rest_context)
		}
	})

	// Access token processing middleware
	e.Use(middleware.AccessTokenMiddleware)

	e.GET("/meta", routes.MetaRoute)
	e.POST("/refresh", routes.RefreshRoute)
	e.GET("/is-valid", routes.IsValidRoute)
	e.POST("/user/login", routes.LoginUserRoute)
	e.POST("/user/register", routes.RegisterUserRoute)
	e.POST("/module/login", routes.LoginModuleRoute)
	e.POST("/module/register", routes.RegisterModuleRoute)

	logger.Fatal(e.Start(fmt.Sprintf(":%d", (*config).Port)))
}
