package types

type Wallet struct {
	Id           int    `json:"id"`
	UserEmail    string `json:"user_email"`
	Currencycode string `json:"currency_code"`
	Balance      float64
}
