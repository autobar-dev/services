package postgres

import (
	"time"

	"github.com/autobar-dev/services/currency/types/errors"
	"github.com/autobar-dev/services/currency/types/inputs"
	"github.com/autobar-dev/services/currency/types/interfaces"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type postgresSupportedCurrenciesStoreRow struct {
	Id        uint32    `db:"id"`
	Code      string    `db:"code"`
	Name      string    `db:"name"`
	Enabled   bool      `db:"enabled"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type PostgresSupportedCurrenciesStore struct {
	logger *interfaces.AppLogger

	database *sqlx.DB
}

func NewPostgresSupportedCurrenciesStore(l *interfaces.AppLogger, db *sqlx.DB) (*PostgresSupportedCurrenciesStore, error) {
	return &PostgresSupportedCurrenciesStore{
		logger:   l,
		database: db,
	}, nil
}

func (s *PostgresSupportedCurrenciesStore) Get(code string) (*interfaces.SupportedCurrenciesStoreRow, error) {
	var r postgresSupportedCurrenciesStoreRow

	err := s.database.QueryRowx(`
		SELECT id, code, name, enabled, created_at, updated_at
		FROM supported_currencies
		WHERE code = $1;
	`, code).StructScan(&r)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.NewDatabaseNotFoundError("specified currency not found")
		}

		return nil, errors.NewDatabaseQueryFailError("could not fetch from supported store")
	}

	ret := interfaces.SupportedCurrenciesStoreRow(r)

	return &ret, nil
}

func (s *PostgresSupportedCurrenciesStore) GetAll() (*[]interfaces.SupportedCurrenciesStoreRow, error) {
	ac, err := s.database.Queryx(`
		SELECT id, code, name, enabled, created_at, updated_at
		FROM supported_currencies
		WHERE enabled = true;
	`)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return &[]interfaces.SupportedCurrenciesStoreRow{}, nil
		}

		return nil, err
	}

	ret := []interfaces.SupportedCurrenciesStoreRow{}

	for ac.Next() {
		var ec postgresSupportedCurrenciesStoreRow

		err := ac.StructScan(&ec)

		if err != nil {
			return nil, err
		}

		ret = append(ret, interfaces.SupportedCurrenciesStoreRow(ec))
	}

	return &ret, nil
}

func (s *PostgresSupportedCurrenciesStore) SetEnabled(currency_code string, enabled bool) error {
	_, err := s.database.Exec(`
		UPDATE supported_currencies
		SET	enabled = $1,
				updated_at = CURRENT_TIMESTAMP
		WHERE code = $2;
	`, enabled, currency_code)

	return err
}

func (s *PostgresSupportedCurrenciesStore) Insert(input *inputs.Currency) error {
	_, err := s.database.Exec(`
		INSERT INTO supported_currencies
		(code, name)
		VALUES ($1, $2);
	`, input.Code, input.Name)

	return err
}

func (s *PostgresSupportedCurrenciesStore) Delete(currency_code string) error {
	_, err := s.database.Exec(`
		DELETE FROM supported_currencies
		WHERE code = $1;
	`, currency_code)

	return err
}
