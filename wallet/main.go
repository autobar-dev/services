package main

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	echo "github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"

	"github.com/autobar-dev/services/wallet/middleware"
	"github.com/autobar-dev/services/wallet/repositories"
	routes "github.com/autobar-dev/services/wallet/routes"
	transaction_routes "github.com/autobar-dev/services/wallet/routes/transaction"
	create_transaction_routes "github.com/autobar-dev/services/wallet/routes/transaction/create"
	wallet_routes "github.com/autobar-dev/services/wallet/routes/wallet"
	"github.com/autobar-dev/services/wallet/types"
	"github.com/autobar-dev/services/wallet/utils"
	authrepository "github.com/autobar-dev/shared-libraries/go/auth-repository"
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

	redis_options, err := redis.ParseURL(config.RedisURL)
	redis_client := redis.NewClient(redis_options)

	e := echo.New()
	e.HideBanner = true

	app_context := &types.AppContext{
		MetaFactors: utils.GetMetaFactors(),
		Config:      config,
		Repositories: &types.Repositories{
			Auth:        authrepository.NewAuthRepository(config.AuthServiceURL, types.MicroserviceName),
			Currency:    repositories.NewCurrencyRepository(config.CurrencyServiceURL),
			Wallet:      repositories.NewWalletRepository(database),
			Transaction: repositories.NewTransactionRepository(database),
			Cache:       repositories.NewCacheRepository(redis_client),
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
