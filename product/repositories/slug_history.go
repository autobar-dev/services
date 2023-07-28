package repositories

import (
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresSlugHistoryEntry struct {
	Id        int32     `db:"id"`
	ProductId string    `db:"product_id"`
	Slug      string    `db:"slug"`
	CreatedAt time.Time `db:"created_at"`
}

type SlugHistoryRepository struct {
	db *sqlx.DB
}

func NewSlugHistoryRepository(db *sqlx.DB) *SlugHistoryRepository {
	return &SlugHistoryRepository{db}
}

func (shr SlugHistoryRepository) Get(slug string) (*PostgresSlugHistoryEntry, error) {
	get_slug_history_entry_query := `
	  SELECT id, product_id, slug, created_at
    FROM slug_history
    WHERE slug=$1;
  `

	row := shr.db.QueryRowx(get_slug_history_entry_query, slug)

	var pshe PostgresSlugHistoryEntry
	err := row.StructScan(&pshe)

	if err != nil {
		fmt.Printf("cannot parse postgres slug history entry: %v\n", err)
		return nil, errors.New("slug history entry failed to be parsed")
	}

	return &pshe, nil
}

func (shr SlugHistoryRepository) GetAllSlugsForProduct(product_id string) (*[]PostgresSlugHistoryEntry, error) {
	get_slug_history_entries_query := `
		SELECT id, product_id, slug, created_at
		FROM slug_history
		WHERE product_id = $1
	  ORDER BY created_at ASC;
	`

	rows, err := shr.db.Queryx(get_slug_history_entries_query, product_id)
	if err != nil {
		return nil, err
	}

	pshes := []PostgresSlugHistoryEntry{}

	for rows.Next() {
		var pshe PostgresSlugHistoryEntry
		err = rows.StructScan(&pshe)

		if err != nil {
			fmt.Printf("cannot parse postgres slug history entry: %v\n", err)
			return nil, errors.New("some slug history entries failed to be parsed")
		}

		pshes = append(pshes, pshe)
	}

	return &pshes, nil
}

func (shr SlugHistoryRepository) Create(product_id string, slug string) error {
	create_slug_history_entry_query := `
	  INSERT INTO slug_history (
			product_id, slug
		) VALUES ($1, $2);
  `

	_, err := shr.db.Exec(create_slug_history_entry_query, product_id, slug)

	return err
}
