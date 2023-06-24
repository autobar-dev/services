package utils

import "go.a5r.dev/services/wallet/types"
import "errors"

type WalletBuilder struct {
	id            int
	user_email    string
	balance       int
	currency_code string
}

func NewWalletBuilder() *WalletBuilder {
	return &WalletBuilder{
		id:            0,
		user_email:    "",
		balance:       0,
		currency_code: "",
	}
}

func (wb *WalletBuilder) Build() (*types.Wallet, error) {
	if wb.id == 0 {
		return nil, errors.New("wallet id not set")
	}

	if wb.currency_code == "" {
		return nil, errors.New("currency code has not been set")
	}

	return &types.Wallet{
		Id:           wb.id,
		UserEmail:    wb.user_email,
		Balance:      wb.balance,
		CurrencyCode: wb.currency_code,
	}, nil
}

func (wb *WalletBuilder) SetId(id int) *WalletBuilder {
	wb.id = id
	return wb
}

func (wb *WalletBuilder) SetUserEmail(email string) *WalletBuilder {
	wb.user_email = email
	return wb
}

func (wb *WalletBuilder) AddTransaction(t types.Transaction) *WalletBuilder {
	switch t.TransactionType {
	case types.TransactionTypeDeposit:
		wb.AddDepositTransaction(t)
	case types.TransactionTypeWithdraw:
		wb.AddWithdrawTransaction(t)
	case types.TransactionTypePurchase:
		wb.AddPurchaseTransaction(t)
	case types.TransactionTypeRefund:
		wb.AddRefundTransaction(t)
	case types.TransactionTypeCurrencyChange:
		wb.AddCurrencyChangeTransaction(t)
	}

	return wb
}

func (wb *WalletBuilder) AddDepositTransaction(t types.Transaction) *WalletBuilder {
	wb.balance += t.Value
	return wb
}

func (wb *WalletBuilder) AddWithdrawTransaction(t types.Transaction) *WalletBuilder {
	wb.balance -= t.Value
	return wb
}

func (wb *WalletBuilder) AddPurchaseTransaction(t types.Transaction) *WalletBuilder {
	wb.balance -= t.Value
	return wb
}

func (wb *WalletBuilder) AddRefundTransaction(t types.Transaction) *WalletBuilder {
	wb.balance += t.Value
	return wb
}

func (wb *WalletBuilder) AddCurrencyChangeTransaction(t types.Transaction) *WalletBuilder {
	wb.balance = t.Value
	wb.currency_code = t.CurrencyCode
	return wb
}
