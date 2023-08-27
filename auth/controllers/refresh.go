package controllers

import (
	"github.com/autobar-dev/services/auth/types"
	"github.com/autobar-dev/services/auth/utils"
	authrepository "github.com/autobar-dev/shared-libraries/go/auth-repository"
)

func Refresh(ac *types.AppContext, refresh_token string) (*types.Tokens, error) {
	ap := ac.Providers.Auth
	aur := ac.Repositories.AuthUser

	owner, err := ap.GetRefreshTokenOwner(refresh_token)
	if err != nil {
		return nil, err
	}

	if owner.Type == authrepository.UserTokenOwnerType {
		auth_user, err := aur.GetById(owner.Identifier)
		if err != nil {
			return nil, err
		}

		access_token := utils.GenerateUserAccessToken(ac.Config.JwtSecret, auth_user.Id, auth_user.Role)

		return &types.Tokens{
			AccessToken:  access_token,
			RefreshToken: refresh_token,
		}, nil
	} else {
		access_token := utils.GenerateModuleAccessToken(ac.Config.JwtSecret, owner.Identifier)

		return &types.Tokens{
			AccessToken:  access_token,
			RefreshToken: refresh_token,
		}, nil

	}
}
