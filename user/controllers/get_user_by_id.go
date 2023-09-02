package controllers

import (
	"fmt"

	"github.com/autobar-dev/services/user/types"
	"github.com/autobar-dev/services/user/utils"
)

func GetUserById(ac *types.AppContext, id string) (*types.UserExtended, error) {
	cr := *ac.Repositories.Cache
	ur := ac.Repositories.User
	wr := ac.Repositories.Wallet

	ru, _ := cr.GetUser(id)
	if ru != nil {
		u := utils.RedisUserToUser(*ru)
		w, err := wr.GetWallet(u.Id)
		if err != nil {
			return nil, err
		}

		return utils.UserToUserExtended(*u, *w), nil
	}

	pu, err := ur.Get(id)
	if err != nil {
		return nil, err
	}

	user := utils.PostgresUserToUser(*pu)
	w, err := wr.GetWallet(user.Id)
	if err != nil {
		return nil, err
	}

	err = cr.SetUser(
		user.Id,
		user.Email,
		user.FirstName,
		user.LastName,
		user.DateOfBirth,
		user.Locale,
		user.IdentityVerificationId,
		user.IdentityVerificationSource,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		fmt.Printf("failed to set cache for user_id->user: %v\n", err)
	}

	return utils.UserToUserExtended(*user, *w), nil
}
