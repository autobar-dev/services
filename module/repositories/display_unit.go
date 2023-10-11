package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresDisplayUnit struct {
	Id                     int32     `db:"id"`
	Symbol                 string    `db:"symbol"`
	DivisorFromMillilitres float64   `db:"divisor_from_millilitres"`
	DecimalsDisplayed      int32     `db:"decimals_displayed"`
	CreatedAt              time.Time `db:"created_at"`
	UpdatedAt              time.Time `db:"updated_at"`
}

type DisplayUnitRepository struct {
	db *sqlx.DB
}

func NewDisplayUnitRepository(db *sqlx.DB) *DisplayUnitRepository {
	return &DisplayUnitRepository{db}
}

func (dur *DisplayUnitRepository) GetDisplayUnit(id int32) (*PostgresDisplayUnit, error) {
	get_display_unit_query := `
		SELECT id, symbol, divisor_from_millilitres, decimals_displayed, created_at, updated_at
		FROM display_units
		WHERE id=$1;
	`

	row := dur.db.QueryRowx(get_display_unit_query, id)

	var pdu PostgresDisplayUnit
	err := row.StructScan(&pdu)
	if err != nil {
		return nil, err
	}

	return &pdu, nil
}
