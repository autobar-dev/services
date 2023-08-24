package types

type RefreshTokenOwnerType string

const (
	UserRefreshTokenOwnerType   RefreshTokenOwnerType = "user"
	ModuleRefreshTokenOwnerType RefreshTokenOwnerType = "module"
)

type RefreshTokenOwner struct {
	Type       RefreshTokenOwnerType
	Identifier string
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
