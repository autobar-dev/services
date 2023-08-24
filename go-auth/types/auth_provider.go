package types

type AuthProvider interface {
	Login(email string, password string, remember_me bool) (*Tokens, error)
	Logout(refresh_token string) error
	Register(email string, password string) (*Tokens, error)
	Refresh(refresh_token string) (*Tokens, error)
	VerifyAccessToken(access_token string) (*AccessTokenPayload, error)
}
