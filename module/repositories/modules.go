package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresModule struct {
	Id           int32  `db:"id"`
	SerialNumber string `db:"serial_number"`

	StationSlug *string `db:"station_slug"`
	ProductSlug *string `db:"product_slug"`

	Prices string `db:"prices"`

	CreatedAt time.Time `db:"created_at"`
}

type ModuleRepository struct {
	db *sqlx.DB
}

func NewModuleRepository(db *sqlx.DB) *ModuleRepository {
	return &ModuleRepository{db}
}

func (mr ModuleRepository) Get(serial_number string) (*PostgresModule, error) {
	get_module_query := `
		SELECT id, serial_number, station_slug, product_slug, prices, created_at
		FROM modules
		WHERE serial_number=$1;
	`

	row := mr.db.QueryRowx(get_module_query, serial_number)

	var pm PostgresModule
	err := row.StructScan(&pm)

	if err != nil {
		return nil, err
	}

	return &pm, nil
}

func (mr ModuleRepository) Create(serial_number string) (*string, error) {
	get_module_query := `
		INSERT INTO modules
		(serial_number)
		VALUES ($1)
		RETURNING serial_number;
	`

	row := mr.db.QueryRowx(get_module_query, serial_number)

	var sn string
	err := row.Scan(&sn)

	if err != nil {
		return nil, err
	}

	return &sn, nil
}
