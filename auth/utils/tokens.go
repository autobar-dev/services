package utils

import (
	"time"

	"github.com/autobar-dev/services/auth/types"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateUserAccessToken(jwt_secret string, user_id string, role string) string {
	expires_at := time.Now().UTC().Add(types.AccessTokenValidDuration)

	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = &types.AccessTokenClaims{
		&jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expires_at),
			Subject:   user_id,
		},
		"user",
		&role,
	}

	signed_token, _ := token.SignedString([]byte(jwt_secret))

	return signed_token
}

func GenerateModuleAccessToken(jwt_secret string, serial_number string) string {
	expires_at := time.Now().UTC().Add(types.AccessTokenValidDuration)

	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = &types.AccessTokenClaims{
		&jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expires_at),
			Subject:   serial_number,
		},
		"module",
		nil,
	}
	signed_token, _ := token.SignedString([]byte(jwt_secret))

	return signed_token
}
