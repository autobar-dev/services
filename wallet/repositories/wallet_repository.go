package repositories

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresWallet struct {
	id            int
	user_email    string
	currency_code string
}

type WalletRepository struct {
	db *sqlx.DB
}

func NewWalletRepository(db *sqlx.DB) *WalletRepository {
	return &WalletRepository{db}
}

func (wr WalletRepository) Get(user_email string) (*PostgresWallet, error) {
	get_wallet_query := `
    SELECT id, user_email, currency_code
    FROM wallets
    WHERE user_email=$1;
  `

	result := wr.db.QueryRowx(get_wallet_query, user_email)

	fmt.Println(result)

	return nil, errors.New("no wallet found for specified user")
}

func (wr WalletRepository) Create(user_email string, currency_code string) error {
	create_wallet_query := `
    INSERT INTO wallets
    (user_email, currency_code)
    VALUES ($1, $2);
  `

	_, err := wr.db.Exec(create_wallet_query, user_email, currency_code)

	return err
}
