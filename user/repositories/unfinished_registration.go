package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresUnfinishedRegistration struct {
	Id                            string    `db:"id"`
	Email                         string    `db:"email"`
	FirstName                     string    `db:"first_name"`
	LastName                      string    `db:"last_name"`
	DateOfBirth                   time.Time `db:"date_of_birth"`
	Locale                        string    `db:"locale"`
	EmailConfirmed                bool      `db:"email_confirmed"`
	EmailConfirmationCode         string    `db:"email_confirmation_code"`
	EmailConfirmationCodeIssuedAt string    `db:"email_confirmation_code_issued_at"`
	CreatedAt                     time.Time `db:"created_at"`
	UpdatedAt                     time.Time `db:"updated_at"`
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

func (urr UnfinishedRegistrationRepository) ChangeEmail(old_email string, new_email string) error {
	change_email_query := `
		UPDATE unfinished_registrations
		SET email = $2
		WHERE email = $1
		RETURNING id;
	`

	r := urr.db.QueryRowx(change_email_query, old_email, new_email)
	return r.Err()
}

func (urr UnfinishedRegistrationRepository) Create(
	email string,
	first_name string,
	last_name string,
	date_of_birth time.Time,
	locale string,
	email_confirmation_code string,
	email_confirmation_code_issued_at time.Time,
) error {
	create_with_personal_details_query := `
		INSERT INTO unfinished_registrations (
			email, first_name, last_name, date_of_birth, locale, email_confirmation_code, email_confirmation_code_issued_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id;
	`

	r := urr.db.QueryRowx(
		create_with_personal_details_query,
		email,
		first_name,
		last_name,
		date_of_birth,
		locale,
		email_confirmation_code,
		email_confirmation_code_issued_at,
	)
	return r.Err()
}

func (urr UnfinishedRegistrationRepository) ConfirmEmail(
	email string,
	email_confirmation_code string,
	current_time time.Time,
) error {
	confirm_email_query := `
		UPDATE unfinished_registrations
		SET email_confirmed = true
		WHERE email = $1;
	`

	r := urr.db.QueryRowx(confirm_email_query, email)
	return r.Err()
}
