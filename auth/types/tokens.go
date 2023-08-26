package types

import "github.com/golang-jwt/jwt/v4"

type TokenOwnerType string

const (
	UserTokenOwnerType   TokenOwnerType = "user"
	ModuleTokenOwnerType TokenOwnerType = "module"
)

type AccessTokenClaims struct {
	*jwt.RegisteredClaims
	ClientType TokenOwnerType `json:"sub_typ"`
	Role       *string        `json:"rol"`
}

type AccessTokenPayload struct {
	ClientType TokenOwnerType
	Identifier string
	Role       *string
}

type RefreshTokenOwner struct {
	Type       TokenOwnerType
	Identifier string
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
