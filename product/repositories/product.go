package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresProduct struct {
	Id           string    `db:"id"`
	Names        string    `db:"names"`
	Descriptions string    `db:"descriptions"`
	Cover        *string   `db:"cover"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (pr ProductRepository) Get(id string) (*PostgresProduct, error) {
	get_product_query := `
		SELECT id, names, descriptions, cover, created_at, updated_at
		FROM products
		WHERE id = $1;
	`

	row := pr.db.QueryRowx(get_product_query, id)

	var pp PostgresProduct
	err := row.StructScan(&pp)

	if err != nil {
		return nil, err
	}

	return &pp, nil
}
