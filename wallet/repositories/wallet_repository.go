package repositories

import (
	"errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresWallet struct {
	Id        int    `db:"id"`
	UserEmail string `db:"user_email"`
}

type WalletRepository struct {
	db *sqlx.DB
}

func NewWalletRepository(db *sqlx.DB) *WalletRepository {
	return &WalletRepository{db}
}

func (wr WalletRepository) Get(user_email string) (*PostgresWallet, error) {
	get_wallet_query := `
    SELECT id, user_email
    FROM wallets
    WHERE user_email=$1;
  `

	result := wr.db.QueryRowx(get_wallet_query, user_email)

	var wallet PostgresWallet
	if err := result.StructScan(&wallet); err != nil {
		return nil, errors.New("wallet not found or database error")
	}

	return &wallet, nil
}

func (wr WalletRepository) Create(user_email string) error {
	create_wallet_query := `
    INSERT INTO wallets
    (user_email)
    VALUES ($1);
  `

	_, err := wr.db.Exec(create_wallet_query, user_email)

	return err
}
