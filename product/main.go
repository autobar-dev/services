package main

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	echo "github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/meilisearch/meilisearch-go"
	"github.com/redis/go-redis/v9"

	"github.com/autobar-dev/services/product/repositories"
	"github.com/autobar-dev/services/product/routes"
	"github.com/autobar-dev/services/product/types"
	"github.com/autobar-dev/services/product/utils"
	filerepository "github.com/autobar-dev/shared-libraries/go/file-repository"
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

	meili_client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   config.MeiliURL,
		APIKey: config.MeiliApiKey,
	})

	e := echo.New()
	e.HideBanner = true

	app_context := &types.AppContext{
		MetaFactors: utils.GetMetaFactors(),
		Config:      config,
		Repositories: &types.Repositories{
			Cache:       repositories.NewCacheRepository(redis_client),
			Product:     repositories.NewProductRepository(database),
			SlugHistory: repositories.NewSlugHistoryRepository(database),
			Meili:       repositories.NewMeiliRepository(meili_client),
			File:        filerepository.NewFileRepository(config.FileServiceURL, types.MicroserviceName),
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
	e.GET("/", routes.GetProductRoute)
	e.POST("/new", routes.CreateProductRoute)
	e.PUT("/edit", routes.EditProductRoute)
	e.GET("/all", routes.GetAllProductsRoute)
	e.POST("/search", routes.SearchProductsRoute)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", (*config).Port)))
}
