package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresUnfinishedRegistration struct {
	Id                             string    `db:"id"`
	Email                          string    `db:"email"`
	Locale                         string    `db:"locale"`
	EmailConfirmationCode          string    `db:"email_confirmation_code"`
	EmailConfirmationCodeExpiresAt time.Time `db:"email_confirmation_code_expires_at"`
	CreatedAt                      time.Time `db:"created_at"`
	UpdatedAt                      time.Time `db:"updated_at"`
}

type UnfinishedRegistrationRepository struct {
	db *sqlx.DB
}

func NewUnfinishedRegistrationRepository(db *sqlx.DB) *UnfinishedRegistrationRepository {
	return &UnfinishedRegistrationRepository{db}
}

func (urr UnfinishedRegistrationRepository) Get(id string) (*PostgresUnfinishedRegistration, error) {
	get_unfinished_registration_query := `
		SELECT *
		FROM unfinished_registrations
		WHERE id = $1;
	`

	row := urr.db.QueryRowx(get_unfinished_registration_query, id)

	var pur PostgresUnfinishedRegistration
	err := row.StructScan(&pur)
	if err != nil {
		return nil, err
	}

	return &pur, nil
}

func (urr UnfinishedRegistrationRepository) GetByEmail(email string) (*PostgresUnfinishedRegistration, error) {
	get_unfinished_registration_query := `
		SELECT *
		FROM unfinished_registrations
		WHERE email = $1;
	`

	row := urr.db.QueryRowx(get_unfinished_registration_query, email)

	var pur PostgresUnfinishedRegistration
	err := row.StructScan(&pur)
	if err != nil {
		return nil, err
	}

	return &pur, nil
}

func (urr UnfinishedRegistrationRepository) GetByConfirmationCode(
	confirmation_code string,
) (*PostgresUnfinishedRegistration, error) {
	get_unfinished_registration_query := `
		SELECT *
		FROM unfinished_registrations
		WHERE email_confirmation_code = $1;
	`

	row := urr.db.QueryRowx(get_unfinished_registration_query, confirmation_code)

	var pur PostgresUnfinishedRegistration
	err := row.StructScan(&pur)
	if err != nil {
		return nil, err
	}

	return &pur, nil
}

func (urr UnfinishedRegistrationRepository) Create(
	id string,
	email string,
	locale string,
	email_confirmation_code string,
	email_confirmation_code_expires_at time.Time,
) error {
	create_user_query := `
		INSERT INTO unfinished_registrations (
			id, email, locale, email_confirmation_code, email_confirmation_code_expires_at
		) VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`

	row := urr.db.QueryRowx(
		create_user_query,
		id,
		email,
		locale,
		email_confirmation_code,
		email_confirmation_code_expires_at,
	)
	var ret_id string
	err := row.Scan(&ret_id)
	if err != nil {
		return err
	}

	return nil
}

func (urr UnfinishedRegistrationRepository) Delete(id string) error {
	delete_unfinished_registration_query := `
		DELETE FROM unfinished_registrations
		WHERE id = $1;
	`

	_, err := urr.db.Exec(delete_unfinished_registration_query, id)
	if err != nil {
		return err
	}

	return nil
}
