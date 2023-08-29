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

	var access_token string

	if owner.Type == authrepository.UserTokenOwnerType {
		auth_user, err := aur.GetById(owner.Identifier)
		if err != nil {
			return nil, err
		}

		access_token = utils.GenerateUserAccessToken(ac.Config.JwtSecret, auth_user.Id, auth_user.Role)
	} else {
		access_token = utils.GenerateModuleAccessToken(ac.Config.JwtSecret, owner.Identifier)
	}

	new_refresh_token, err := ap.UpdateRefreshToken(refresh_token)
	if err != nil {
		return nil, err
	}

	return &types.Tokens{
		AccessToken:  access_token,
		RefreshToken: *new_refresh_token,
	}, nil
}
