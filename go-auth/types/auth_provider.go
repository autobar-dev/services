package types

type AuthProvider interface {
	LoginUser(email string, password string, remember_me bool) (*string, error)
	RegisterUser(user_id string, email string, password string) (*string, error)
	LoginModule(serial_number string, private_key string) (*string, error)
	RegisterModule(serial_number string, private_key string) (*string, error)
	GetRefreshTokenOwner(refresh_token string) (*RefreshTokenOwner, error)
	InvalidateRefreshTokenById(token_id string) error
	InvalidateRefreshTokenByToken(refresh_token string) error
}
