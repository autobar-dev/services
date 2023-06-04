package main

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	echo "github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"go.a5r.dev/services/wallet/repositories"
	wallet_routes "go.a5r.dev/services/wallet/routes/wallet"
	"go.a5r.dev/services/wallet/types"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("failed to load .env file", err)
		os.Exit(1)
	}

	config, err := types.LoadEnvVars()
	if err != nil {
		fmt.Println("failed to load .env file", err)
		os.Exit(1)
	}

	database, err := sqlx.Connect("postgres", config.DatabaseURL)
	if err != nil {
		fmt.Println("failed to connect to database", err)
		os.Exit(1)
	}

	e := echo.New()
	e.HideBanner = true

	app_context := &types.AppContext{
		Message:  "yooo sup",
		Database: database,
		Repositories: &types.Repositories{
			Wallet: repositories.NewWalletRepository(database),
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

	e.POST("/wallet/create", wallet_routes.CreateRoute)
	e.GET("/wallet", wallet_routes.GetRoute)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", (*config).Port)))
}
