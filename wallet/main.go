package main

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	echo "github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"go.a5r.dev/services/wallet/repositories"
	routes "go.a5r.dev/services/wallet/routes"
	transaction_routes "go.a5r.dev/services/wallet/routes/transaction"
	create_transaction_routes "go.a5r.dev/services/wallet/routes/transaction/create"
	wallet_routes "go.a5r.dev/services/wallet/routes/wallet"
	"go.a5r.dev/services/wallet/types"
	"go.a5r.dev/services/wallet/utils"
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

	e := echo.New()
	e.HideBanner = true

	app_context := &types.AppContext{
		Meta:     utils.LoadMeta(),
		Database: database,
		Repositories: &types.Repositories{
			Currency:    repositories.NewCurrencyRepository(config.CurrencyServiceURL),
			Wallet:      repositories.NewWalletRepository(database),
			Transaction: repositories.NewTransactionRepository(database),
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
	e.GET("/wallet/", wallet_routes.GetRoute)
	e.POST("/wallet/create", wallet_routes.CreateRoute)
	e.GET("/transaction/get", transaction_routes.GetRoute)
	e.GET("/transaction/get-all", transaction_routes.GetAllRoute)
	e.POST("/transaction/create/deposit", create_transaction_routes.DepositRoute)
	e.POST("/transaction/create/withdraw", create_transaction_routes.WithdrawRoute)
	e.POST("/transaction/create/purchase", create_transaction_routes.PurchaseRoute)
	e.POST("/transaction/create/refund", create_transaction_routes.RefundRoute)
	e.POST("/transaction/create/currency-change", create_transaction_routes.CurrencyChangeRoute)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", (*config).Port)))
}
