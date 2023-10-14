package controllers

import (
	"github.com/autobar-dev/services/module/types"
	"github.com/autobar-dev/services/module/utils"
)

func GetActivationSession(
	app_context *types.AppContext,
	user_id string,
) (*types.ActivationSession, error) {
	sr := app_context.Repositories.State

	as_id, err := sr.GetActivationSessionIdForUser(user_id)
	if err != nil {
		return nil, err
	}

	ras, err := sr.GetActivationSession(*as_id)
	if err != nil {
		return nil, err
	}

	return utils.RedisActivationSessionToActivationSession(*ras), nil
}
