package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresUser struct {
	Id                         string    `db:"id"`
	Email                      string    `db:"email"`
	FirstName                  string    `db:"first_name"`
	LastName                   string    `db:"last_name"`
	DateOfBirth                time.Time `db:"date_of_birth"`
	Locale                     string    `db:"locale"`
	IdentityVerificationId     *string   `db:"identity_verification_id"`
	IdentityVerificationSource *string   `db:"identity_verification_source"`
	CreatedAt                  time.Time `db:"created_at"`
	UpdatedAt                  time.Time `db:"updated_at"`
}

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur UserRepository) Get(id string) (*PostgresUser, error) {
	get_user_query := `
		SELECT *
		FROM users
		WHERE id = $1;
	`

	row := ur.db.QueryRowx(get_user_query, id)

	var pu PostgresUser
	err := row.StructScan(&pu)
	if err != nil {
		return nil, err
	}

	return &pu, nil
}

func (ur UserRepository) GetByEmail(email string) (*PostgresUser, error) {
	get_user_query := `
		SELECT *
		FROM users
		WHERE email = $1;
	`

	row := ur.db.QueryRowx(get_user_query, email)

	var pu PostgresUser
	err := row.StructScan(&pu)
	if err != nil {
		return nil, err
	}

	return &pu, nil
}

func (ur UserRepository) Create(
	id string,
	email string,
	first_name string,
	last_name string,
	date_of_birth time.Time,
	locale string,
) error {
	create_user_query := `
		INSERT INTO users (
			id, email, first_name, last_name, date_of_birth, locale
		) VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;
	`

	result := ur.db.QueryRowx(create_user_query, id, email, first_name, last_name, date_of_birth, locale)

	var _user_id string
	err := result.Scan(&_user_id)
	if err != nil {
		return err
	}

	return nil
}
