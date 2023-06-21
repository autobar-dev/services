package types

type Wallet struct {
	Id           int    `json:"id"`
	UserEmail    string `json:"user_email"`
	CurrencyCode string `json:"currency_code"`
	Balance      int    `json:"balance"`
}
