package main

import (
	"fmt"
	"os"

	"github.com/autobar-dev/services/currency/stores/postgres"
	"github.com/autobar-dev/services/currency/types"
	"github.com/autobar-dev/services/currency/types/interfaces"
	"github.com/charmbracelet/log"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	var l interfaces.AppLogger = log.Default()

	// Load env file
	err := godotenv.Load(".env")

	if err != nil {
		l.Error("Not able to load .env file")
	}

	cs := os.Getenv("DB_CONNECTION_STRING")
	db, err := sqlx.Connect("postgres", cs)

	if err != nil {
		l.Error("Error connecting to the database.")
		os.Exit(1)
	}

	// Initialize rate store
	rs, err := postgres.NewPostgresRateStore(&l, db)

	if err != nil {
		l.Error(err)
		os.Exit(1)
	}

	l.Info("Connected to the database.")

	// Initialize supported currencies store
	scs, err := postgres.NewPostgresSupportedCurrenciesStore(&l, db)

	if err != nil {
		l.Error(err)
		os.Exit(1)
	}

	// Initialize stores
	stores := &types.AppStores{
		RateStore:                rs,
		SupportedCurrenciesStore: scs,
	}

	r, err := stores.RateStore.GetRate("EUR", "UAH")
	fmt.Printf("rate=%#v\nerr=%#v\n\n", r, err)

	av, err := stores.SupportedCurrenciesStore.IsSupported("UAH")
	fmt.Printf("sc=%#v\nerr=%#v\n\n", av, err)

	ec, err := stores.SupportedCurrenciesStore.GetAll()
	fmt.Printf("ec=%#v\nerr=%#v\n\n", ec, err)

	// r, err := prs.GetRate("EUR", "PLN")

	// if err != nil {
	// 	l.Error("Error getting rate:", err)
	// 	os.Exit(1)
	// }

	// l.Info(fmt.Sprintf("EUR->PLN = %f, updated at %v, with id %d", r.Rate, r.UpdatedAt, r.Id))

	// Initialize supported currencies store
	// scs, err := postgres.New
}
