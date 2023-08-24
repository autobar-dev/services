package types

import "time"

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AccessTokenPayload struct {
	UserId  string
	Role    string
	Expires time.Time
}
