package controllers

import (
	"github.com/autobar-dev/services/auth/types"
)

func Logout(ac *types.AppContext, refresh_token string) error {
	ap := ac.Providers.Auth

	err := ap.InvalidateRefreshTokenByToken(refresh_token)
	if err != nil {
		return err
	}

	return nil
}
