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

	rand.Seed(time.Now().UnixNano())

	e := echo.New()
	e.HideBanner = true

	app_context := &types.AppContext{
		Meta: utils.LoadMeta(),
		Repositories: &types.Repositories{
			Module: repositories.NewModuleRepository(database),
			Auth:   repositories.NewAuthRepository(config.AuthServiceURL),
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
	e.POST("/create", routes.CreateModuleRoute)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", (*config).Port)))
}
