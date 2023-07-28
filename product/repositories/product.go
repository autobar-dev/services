package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresProduct struct {
	Id           string    `db:"id"`
	Names        string    `db:"names"`
	Descriptions string    `db:"descriptions"`
	Cover        string    `db:"cover"`
	Enabled      bool      `db:"enabled"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type PostgresEditProductInput struct {
	Names        *map[string]string `json:"names"`
	Descriptions *map[string]string `json:"descriptions"`
	Cover        *string            `json:"cover"`
	Enabled      *bool              `json:"enabled"`
}

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (pr ProductRepository) Get(id string) (*PostgresProduct, error) {
	get_product_query := `
		SELECT id, names, descriptions, cover, enabled, created_at, updated_at
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

func (pr ProductRepository) GetAll() (*[]PostgresProduct, error) {
	get_all_products_query := `
		SELECT id, names, descriptions, cover, enabled, created_at, updated_at
		FROM products;
	`

	rows, err := pr.db.Queryx(get_all_products_query)
	if err != nil {
		return nil, err
	}

	products := []PostgresProduct{}

	for rows.Next() {
		var pp PostgresProduct

		err = rows.StructScan(&pp)
		if err != nil {
			fmt.Printf("cannot parse postgres product: %v\n", err)
			return nil, errors.New("some products failed to be parsed")
		}

		products = append(products, pp)
	}

	return &products, nil
}

func (pr ProductRepository) Create(names map[string]string, descriptions map[string]string, cover *string, enabled bool) (*string, error) {
	names_bytes, _ := json.Marshal(names)
	names_str := string(names_bytes)

	descriptions_bytes, _ := json.Marshal(descriptions)
	descriptions_str := string(descriptions_bytes)

	create_product_query := `
		INSERT INTO products (
			names, descriptions, cover, enabled
		) VALUES ($1, $2, $3, $4)
		RETURNING id;
	`

	result := pr.db.QueryRowx(create_product_query, names_str, descriptions_str, cover, enabled)

	var product_id string

	err := result.Scan(&product_id)
	if err != nil {
		return nil, err
	}

	return &product_id, nil
}

func (pr ProductRepository) Edit(id string, input *PostgresEditProductInput) error {
	arg_counter := 2
	args_list := make([]interface{}, 0)
	args_names_list := []string{}

	args_list = append(args_list, id) // first argument is always the id

	if input.Names != nil {
		names_bytes, _ := json.Marshal(input.Names)
		names_str := string(names_bytes)

		args_names_list = append(args_names_list, fmt.Sprintf("names=$%d", arg_counter))
		args_list = append(args_list, names_str)
		arg_counter += 1
	}
	if input.Descriptions != nil {
		descriptions_bytes, _ := json.Marshal(input.Descriptions)
		descriptions_str := string(descriptions_bytes)

		args_names_list = append(args_names_list, fmt.Sprintf("descriptions=$%d", arg_counter))
		args_list = append(args_list, descriptions_str)
		arg_counter += 1
	}
	if input.Cover != nil {
		args_names_list = append(args_names_list, fmt.Sprintf("cover=$%d", arg_counter))
		args_list = append(args_list, input.Cover)
		arg_counter += 1
	}
	if input.Enabled != nil {
		args_names_list = append(args_names_list, fmt.Sprintf("enabled=$%d", arg_counter))
		args_list = append(args_list, input.Enabled)
		arg_counter += 1
	}

	edit_product_query := `
		UPDATE products SET `
	edit_product_query += strings.Join(args_names_list, ", ")
	edit_product_query += " WHERE id=$1;"

	fmt.Printf("update products query: %s\n", edit_product_query)
	fmt.Printf("update products arguments: %+v\n", args_list)

	_, err := pr.db.Exec(edit_product_query, args_list...)
	return err
}
