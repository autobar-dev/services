package postgres

import (
	"time"

	"github.com/autobar-dev/services/currency/types/errors"
	"github.com/autobar-dev/services/currency/types/interfaces"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresSupportedCurrenciesStoreRow struct {
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
	var r PostgresSupportedCurrenciesStoreRow

	err := s.database.QueryRowx(`
		SELECT id, code, name, enabled, created_at, updated_at
		FROM supported_currencies
		WHERE code = $1;
	`, code).StructScan(&r)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}

		return nil, errors.NewDatabaseQueryFailError("Could not fetch from supported store.")
	}

	ret := interfaces.SupportedCurrenciesStoreRow(r)

	return &ret, nil
}

func (s *PostgresSupportedCurrenciesStore) IsSupported(code string) (bool, error) {
	sc, err := s.Get(code)

	if err != nil {
		return false, err
	}

	if sc == nil {
		return false, nil
	}

	return sc.Enabled, nil
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
		var ec PostgresSupportedCurrenciesStoreRow

		err := ac.StructScan(&ec)

		if err != nil {
			return nil, err
		}

		ret = append(ret, interfaces.SupportedCurrenciesStoreRow(ec))
	}

	return &ret, nil
}
