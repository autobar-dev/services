package repositories

import (
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type PostgresTransactionType string

const (
	PostgresTransactionTypeDeposit  PostgresTransactionType = "deposit"
	PostgresTransactionTypeWithdraw PostgresTransactionType = "withdraw"
	PostgresTransactionTypePurchase PostgresTransactionType = "purchase"
	PostgresTransactionTypeRefund   PostgresTransactionType = "refund"
)

type PostgresTransaction struct {
	Id              string                  `db:"id"`
	WalletId        int                     `db:"wallet_id"`
	TransactionType PostgresTransactionType `db:"type"`
	Value           int                     `db:"value"`
	CurrencyCode    string                  `db:"currency_code"`
	CreatedAt       time.Time               `db:"created_at"`
}

type TransactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{db}
}

func (tr TransactionRepository) GetForWallet(wallet_id int) (*[]PostgresTransaction, error) {
	get_transactions_for_wallet_query := `
    SELECT id, wallet_id, type, value, currency_code, created_at
    FROM transactions
    WHERE wallet_id=$1
		ORDER BY created_at ASC;
  `

	rows, err := tr.db.Queryx(get_transactions_for_wallet_query, wallet_id)

	if err != nil {
		return nil, err
	}

	transactions := []PostgresTransaction{}

	for rows.Next() {
		var pt PostgresTransaction
		err = rows.StructScan(&pt)

		if err != nil {
			fmt.Printf("cannot parse postgres transaction: %v\n", err)
			return nil, errors.New("some transactions failed to be parsed")
		}

		transactions = append(transactions, pt)
	}

	return &transactions, nil
}
