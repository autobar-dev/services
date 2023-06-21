package types

import "time"

type TransactionType string

const (
	TransactionTypeDeposit        TransactionType = "deposit"
	TransactionTypeWithdraw       TransactionType = "withdraw"
	TransactionTypePurchase       TransactionType = "purchase"
	TransactionTypeRefund         TransactionType = "refund"
	TransactionTypeCurrencyChange TransactionType = "currency_change"
)

type Transaction struct {
	Id              string          `json:"id"`
	WalletId        int             `json:"wallet_id"`
	TransactionType TransactionType `json:"type"`
	Value           int             `json:"value"`
	CurrencyCode    string          `json:"currency_code"`
	CreatedAt       time.Time       `json:"created_at"`
}
