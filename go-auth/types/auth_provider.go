package types

type AuthProvider interface {
	LoginUser(email string, password string, remember_me bool) (refresh_token *string, err error)
	RegisterUser(user_id string, email string, password string) error
	LoginModule(serial_number string, private_key string) (refresh_token *string, err error)
	RegisterModule(serial_number string, private_key string) error
	InvalidateRefreshTokenById(token_id string) error
	InvalidateRefreshTokenByToken(refresh_token string) error
	GetRefreshTokenOwner(refresh_token string) (owner *RefreshTokenOwner, err error)
}
