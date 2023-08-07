package types

type Wallet struct {
	Id           int    `json:"id"`
	UserId       string `json:"user_id"`
	CurrencyCode string `json:"currency_code"`
	Balance      int    `json:"balance"`
}
