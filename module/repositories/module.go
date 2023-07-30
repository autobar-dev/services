package repositories

import (
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresModule struct {
	Id           int32  `db:"id"`
	SerialNumber string `db:"serial_number"`

	StationId *string `db:"station_id"`
	ProductId *string `db:"product_id"`

	Enabled bool `db:"enabled"`

	Prices string `db:"prices"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type ModuleRepository struct {
	db *sqlx.DB
}

func NewModuleRepository(db *sqlx.DB) *ModuleRepository {
	return &ModuleRepository{db}
}

func (mr ModuleRepository) Get(serial_number string) (*PostgresModule, error) {
	get_module_query := `
		SELECT id, serial_number, station_id, product_id, enabled, prices, created_at, updated_at
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

func (mr ModuleRepository) GetAll() (*[]PostgresModule, error) {
	get_modules_query := `
		SELECT id, serial_number, station_id, product_id, enabled, prices, created_at, updated_at
		FROM modules;
	`

	rows, err := mr.db.Queryx(get_modules_query)
	if err != nil {
		return nil, err
	}

	modules := []PostgresModule{}

	for rows.Next() {
		var pm PostgresModule
		err = rows.StructScan(&pm)

		if err != nil {
			fmt.Printf("cannot parse postgres module: %v\n", err)
			return nil, errors.New("some modules failed to be parsed")
		}

		modules = append(modules, pm)
	}

	return &modules, nil
}

func (mr ModuleRepository) GetAllForStation(station_id string) (*[]PostgresModule, error) {
	get_modules_for_station_query := `
		SELECT id, serial_number, station_id, product_id, enabled, prices, created_at, updated_at
		FROM modules
		WHERE station_id = $1;
	`

	rows, err := mr.db.Queryx(get_modules_for_station_query, station_id)
	if err != nil {
		return nil, err
	}

	modules := []PostgresModule{}

	for rows.Next() {
		var pm PostgresModule
		err = rows.StructScan(&pm)

		if err != nil {
			fmt.Printf("cannot parse postgres module: %v\n", err)
			return nil, errors.New("some modules failed to be parsed")
		}

		modules = append(modules, pm)
	}

	return &modules, nil
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
