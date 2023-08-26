package controllers

import (
	"github.com/autobar-dev/services/auth/types"
	"github.com/autobar-dev/services/auth/utils"
)

func LoginModule(ac *types.AppContext, serial_number string, private_key string) (*types.Tokens, error) {
	ap := ac.Providers.Auth
	amr := ac.Repositories.AuthModule

	refresh_token, err := ap.LoginModule(serial_number, private_key)
	if err != nil {
		return nil, err
	}

	auth_module, err := amr.GetBySerialNumber(serial_number)
	if err != nil {
		return nil, err
	}

	access_token := utils.GenerateModuleAccessToken(ac.Config.JwtSecret, auth_module.SerialNumber)

	return &types.Tokens{
		AccessToken:  access_token,
		RefreshToken: *refresh_token,
	}, nil
}
