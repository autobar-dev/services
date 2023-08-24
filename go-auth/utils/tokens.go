package utils

import (
	"time"

	"github.com/autobar-dev/services/auth/types"
	"github.com/golang-jwt/jwt"
)

func GenerateUserAccessToken(jwt_secret string, user_id string, role string) string {
	expires_at := time.Now().UTC().Add(types.AccessTokenValidDuration).Format(time.RFC3339)

	jwt_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub_typ": "user",
		"sub":     user_id,
		"rol":     role,
		"exa":     expires_at,
	})
	signed_token, _ := jwt_token.SignedString([]byte(jwt_secret))

	return signed_token
}

func GenerateModuleAccessToken(jwt_secret string, serial_number string) string {
	expires_at := time.Now().UTC().Add(types.AccessTokenValidDuration).Format(time.RFC3339)

	jwt_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub_typ": "module",
		"sub":     serial_number,
		"rol":     nil,
		"exa":     expires_at,
	})
	signed_token, _ := jwt_token.SignedString([]byte(jwt_secret))

	return signed_token
}
