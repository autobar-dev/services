package types

import "time"

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type AccessTokenPayload struct {
	UserId  string
	Scope   []string
	Expires time.Time
}

type AuthProvider interface {
	Login
	Refresh(refresh_token string) *Tokens
	VerifyAccessToken(access_token string) *AccessTokenPayload
}
