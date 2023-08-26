package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type PostgresAuthUser struct {
	Id        string    `db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Role      string    `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type AuthUserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *AuthUserRepository {
	return &AuthUserRepository{db}
}

func (pur *AuthUserRepository) Create(user_id string, email string, password string) error {
	create_user_query := `
		INSERT INTO users (id, email, password)
		VALUES ($1, $2, $3);
	`

	_, err := pur.db.Exec(create_user_query, user_id, email, password)
	if err != nil {
		return err
	}

	return nil
}

func (pur *AuthUserRepository) GetById(user_id string) (*PostgresAuthUser, error) {
	get_user_query := `
 		SELECT *
		FROM users
		WHERE id = $1;
	`

	row := pur.db.QueryRowx(get_user_query, user_id)

	var pu PostgresAuthUser
	err := row.StructScan(&pu)
	if err != nil {
		return nil, err
	}

	return &pu, nil
}

func (pur *AuthUserRepository) GetByEmail(email string) (*PostgresAuthUser, error) {
	get_user_query := `
 		SELECT *
		FROM users
		WHERE email = $1;
	`

	row := pur.db.QueryRowx(get_user_query, email)

	var pu PostgresAuthUser
	err := row.StructScan(&pu)
	if err != nil {
		return nil, err
	}

	return &pu, nil
}
