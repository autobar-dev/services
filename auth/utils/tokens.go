package utils

import (
	"errors"
	"fmt"
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

func ParseAccessToken(jwt_secret string, access_token string) (*types.AccessTokenPayload, error) {
	token, err := jwt.Parse(access_token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprintf("unexpected signing method: %v", token.Header["alg"]))
		}

		return []byte(jwt_secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	identifier, ok := claims["sub"].(string)
	if !ok {
		return nil, errors.New("invalid token")
	}

	var client_type types.TokenOwnerType
	token_client_type, ok := claims["sub_typ"].(string)
	if token_client_type == string(types.UserTokenOwnerType) {
		client_type = types.UserTokenOwnerType
	} else if token_client_type == string(types.ModuleTokenOwnerType) {
		client_type = types.ModuleTokenOwnerType
	} else {
		return nil, errors.New("invalid token")
	}

	var role *string = nil
	token_role, ok := claims["rol"].(string)
	if ok {
		role = &token_role
	}

	return &types.AccessTokenPayload{
		Identifier: identifier,
		ClientType: client_type,
		Role:       role,
	}, nil
}
