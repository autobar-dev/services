package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type PostgresModule struct {
	Id           int       `db:"id"`
	SerialNumber string    `db:"serial_number"`
	PrivateKey   string    `db:"private_key"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type PostgresModuleRepository struct {
	db *sqlx.DB
}

func NewPostgresModuleRepository(db *sqlx.DB) *PostgresModuleRepository {
	return &PostgresModuleRepository{db}
}

func (pmr *PostgresModuleRepository) Create(serial_number string, private_key string) error {
	return nil
}
