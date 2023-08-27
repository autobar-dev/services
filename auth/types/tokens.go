package types

import (
	authrepository "github.com/autobar-dev/shared-libraries/go/auth-repository"
	"github.com/golang-jwt/jwt/v4"
)

type AccessTokenClaims struct {
	*jwt.RegisteredClaims
	ClientType authrepository.TokenOwnerType `json:"sub_typ"`
	Role       *string                       `json:"rol"`
}

type RefreshTokenOwner struct {
	Type       authrepository.TokenOwnerType
	Identifier string
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
