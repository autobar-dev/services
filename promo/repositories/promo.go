package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresPromo struct {
	Id        string    `db:"id"`
	Extension string    `db:"extension"`
	CreatedAt time.Time `db:"created_at"`
}

type FileRepository struct {
	db *sqlx.DB
}

func NewFileRepository(db *sqlx.DB) *FileRepository {
	return &FileRepository{db}
}

func (fr *FileRepository) Get(id string) (*PostgresFile, error) {
	get_file_query := `
		SELECT id, extension, created_at
		FROM files
		WHERE id = $1;
	`

	row := fr.db.QueryRowx(get_file_query, id)

	var pf PostgresFile
	err := row.StructScan(&pf)
	if err != nil {
		return nil, err
	}

	return &pf, nil
}

func (fr *FileRepository) Create(
	id string,
	extension string,
) error {
	create_file_query := `
		INSERT INTO files (
			id, extension
		) VALUES ($1, $2)
		RETURNING id;
	`

	result := fr.db.QueryRowx(create_file_query, id, extension)

	var result_id string

	err := result.Scan(&result_id)
	if err != nil {
		return err
	}

	return nil
}

func (fr *FileRepository) Delete(id string) error {
	delete_file_query := `
		DELETE FROM files
		WHERE id = $1;
	`

	_, err := fr.db.Exec(delete_file_query, id)
	return err
}
