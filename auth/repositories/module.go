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
	create_module_query := `
		INSERT INTO modules (serial_number, private_key)
		VALUES ($1, $2);
	`

	_, err := pmr.db.Exec(create_module_query, serial_number, private_key)
	if err != nil {
		return err
	}

	return nil
}

func (pmr *AuthModuleRepository) GetBySerialNumber(serial_number string) (*PostgresAuthModule, error) {
	get_module_query := `
 		SELECT *
		FROM modules
		WHERE serial_number = $1;
	`

	row := pmr.db.QueryRowx(get_module_query, serial_number)

	var pm PostgresAuthModule
	err := row.StructScan(&pm)
	if err != nil {
		return nil, err
	}

	return &pm, nil
}
