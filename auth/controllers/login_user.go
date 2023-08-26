package controllers

import (
	"fmt"

	"github.com/autobar-dev/services/auth/types"
	"github.com/autobar-dev/services/auth/utils"
)

func LoginUser(ac *types.AppContext, email string, password string, remember_me bool) (*types.Tokens, error) {
	ap := ac.Providers.Auth
	aur := ac.Repositories.AuthUser

	fmt.Println("login user controller")

	refresh_token, err := ap.LoginUser(email, password, remember_me)
	if err != nil {
		return nil, err
	}

	auth_user, err := aur.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	access_token := utils.GenerateUserAccessToken(ac.Config.JwtSecret, auth_user.Id, auth_user.Role)

	return &types.Tokens{
		AccessToken:  access_token,
		RefreshToken: *refresh_token,
	}, nil
}
