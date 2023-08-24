package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type PostgresAuthModule struct {
	Id           int       `db:"id"`
	SerialNumber string    `db:"serial_number"`
	PrivateKey   string    `db:"private_key"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type AuthModuleRepository struct {
	db *sqlx.DB
}

func NewModuleRepository(db *sqlx.DB) *AuthModuleRepository {
	return &AuthModuleRepository{db}
}

func (pmr *AuthModuleRepository) Create(serial_number string, private_key string) error {
	return nil
}

func (pmr *AuthModuleRepository) GetBySerialNumber(serial_number string) (*PostgresAuthModule, error) {
	return nil, nil
}
