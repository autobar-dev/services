package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type PostgresRefreshToken struct {
	Id                 string    `db:"id"`
	UserId             *string   `db:"user_id"`
	ModuleSerialNumber *string   `db:"module_serial_number"`
	Token              string    `db:"token"`
	ExpiresAt          time.Time `db:"expires_at"`
	CreatedAt          time.Time `db:"created_at"`
}

type PostgresRefreshTokenRepository struct {
	db *sqlx.DB
}

func NewPostgresRefreshTokenRepository(db *sqlx.DB) *PostgresRefreshTokenRepository {
	return &PostgresRefreshTokenRepository{db}
}

func (rtr *PostgresRefreshTokenRepository) Create(user_id string, token_value string, valid_for time.Duration) error {
	return nil
}
