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

type RefreshTokenRepository struct {
	db *sqlx.DB
}

func NewRefreshTokenRepository(db *sqlx.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{db}
}

func (rtr *RefreshTokenRepository) CreateForUser(user_id string, token_value string, valid_until time.Time) error {
	create_token_query := `
		INSERT INTO refresh_tokens
		(user_id, token, expires_at)
		VALUES ($1, $2, $3);
	`

	_, err := rtr.db.Exec(create_token_query, user_id, token_value, valid_until)
	if err != nil {
		return err
	}

	return nil
}

func (rtr *RefreshTokenRepository) CreateForModule(serial_number string, token_value string, valid_until time.Time) error {
	create_token_query := `
		INSERT INTO refresh_tokens
		(module_serial_number, token, expires_at)
		VALUES ($1, $2, $3);
	`

	_, err := rtr.db.Exec(create_token_query, serial_number, token_value, valid_until)
	if err != nil {
		return err
	}

	return nil
}

func (rtr *RefreshTokenRepository) GetByToken(token string) (*PostgresRefreshToken, error) {
	get_token_query := `
		SELECT *
		FROM refresh_tokens
		WHERE token = $1;
	`

	row := rtr.db.QueryRowx(get_token_query, token)

	var prt PostgresRefreshToken
	err := row.StructScan(&prt)
	if err != nil {
		return nil, err
	}

	return &prt, nil
}

func (rtr *RefreshTokenRepository) GetById(token_id string) (*PostgresRefreshToken, error) {
	get_token_query := `
		SELECT *
		FROM refresh_tokens
		WHERE id = $1;
	`

	row := rtr.db.QueryRowx(get_token_query, token_id)

	var prt PostgresRefreshToken
	err := row.StructScan(&prt)
	if err != nil {
		return nil, err
	}

	return &prt, nil
}

func (rtr *RefreshTokenRepository) EditByToken(
	token string,
	new_token_value string,
	new_expires_at time.Time,
) error {
	edit_token_query := `
		UPDATE refresh_tokens
		SET token = $1, expires_at = $2
		WHERE token = $3;
	`

	_, err := rtr.db.Exec(edit_token_query, new_token_value, new_expires_at, token)
	return err
}

func (rtr *RefreshTokenRepository) DeleteById(token_id string) error {
	delete_token_query := `
		DELETE FROM refresh_tokens
		WHERE id = $1;
	`

	_, err := rtr.db.Exec(delete_token_query, token_id)
	if err != nil {
		return err
	}

	return nil
}

func (rtr *RefreshTokenRepository) DeleteByToken(token string) error {
	delete_token_query := `
		DELETE FROM refresh_tokens
		WHERE token = $1;
	`

	_, err := rtr.db.Exec(delete_token_query, token)
	if err != nil {
		return err
	}

	return nil
}
