package controllers

import (
	"fmt"

	"github.com/autobar-dev/services/user/types"
	"github.com/autobar-dev/services/user/utils"
)

func GetUserById(ac *types.AppContext, id string) (*types.User, error) {
	cr := *ac.Repositories.Cache
	ur := *ac.Repositories.User

	ru, _ := cr.GetUser(id)
	if ru != nil {
		return utils.RedisUserToUser(*ru), nil
	}

	pu, err := ur.Get(id)
	if err != nil {
		return nil, err
	}

	user := utils.PostgresUserToUser(*pu)

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

	return user, nil
}
