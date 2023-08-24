package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type PostgresUser struct {
	Id        string    `db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type PostgresUserRepository struct {
	db *sqlx.DB
}

func NewPostgresUserRepository(db *sqlx.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db}
}

func (pur *PostgresUserRepository) Create(email string, password string) error {
	return nil
}
