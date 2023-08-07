package types

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type AuthProvider interface {
	Refresh(refresh_token string) *Tokens
	Verify
}
