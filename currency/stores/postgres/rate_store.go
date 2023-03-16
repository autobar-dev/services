package postgres

import (
	"fmt"
	"time"

	"github.com/autobar-dev/services/currency/types/errors"
	"github.com/autobar-dev/services/currency/types/interfaces"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresRateStoreRow struct {
	Id                  uint32    `db:"id"`
	BaseCurrency        string    `db:"base_currency"`
	DestinationCurrency string    `db:"destination_currency"`
	Rate                float32   `db:"rate"`
	UpdatedAt           time.Time `db:"updated_at"`
}

type PostgresRateStore struct {
	logger *interfaces.AppLogger

	database *sqlx.DB
}

func NewPostgresRateStore(l *interfaces.AppLogger, db *sqlx.DB) (*PostgresRateStore, error) {
	return &PostgresRateStore{
		logger:   l,
		database: db,
	}, nil
}

func (s *PostgresRateStore) GetRate(base string, destination string) (*interfaces.RateStoreRow, error) {
	var r PostgresRateStoreRow

	err := s.database.QueryRowx(`
		SELECT rates.id, bsc.code AS base_currency, dsc.code AS destination_currency, rates.rate, rates.updated_at
		FROM rates
		INNER JOIN supported_currencies bsc ON (rates.base_currency_id = bsc.id)
		INNER JOIN supported_currencies dsc ON (rates.destination_currency_id = dsc.id) 
		WHERE bsc.code = $1 AND dsc.code = $2;
	`, base, destination).StructScan(&r)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}

		return nil, errors.NewDatabaseQueryFailError("Failed to fetch rate.")
	}

	ret := interfaces.RateStoreRow(r)

	return &ret, nil
}

func (s *PostgresRateStore) Upsert(base string, destination string, rate float32) error {
	row, _ := s.GetRate(base, destination)
	l := *s.logger

	var query string

	if row == nil {
		l.Info(fmt.Sprintf("Creating new rate entry for %s->%s=%f", base, destination, rate))

		query = fmt.Sprintf(`
			INSERT INTO rates (base_currency, destination_currency, rate)
			VALUES ('%s', '%s', %f);
		`, base, destination, rate)
	} else {
		l.Info(fmt.Sprintf("Updating rate entry for %s->%s=%f", base, destination, rate))

		query = fmt.Sprintf(`
			UPDATE rates
			SET	rate=%f,
					updated_at=CURRENT_TIMESTAMP
			WHERE id=%d;
		`, rate, row.Id)
	}

	_, err := s.database.Exec(query)

	return err
}

func (s *PostgresRateStore) Delete(base string, destination string) error {
	_, err := s.database.Exec(`
		DELETE FROM rates
		WHERE base_currency=$1 AND destination_currency=$2;
	`, base, destination)

	return err
}
