package controllers

import (
	"go.a5r.dev/services/module/types"
	"go.a5r.dev/services/module/utils"
)

func VerifySessionController(app_context *types.AppContext, session_id string) (*types.SessionData, error) {
	ar := app_context.Repositories.Auth

	svs, err := ar.VerifySession(session_id)
	if err != nil {
		return nil, err
	}

	session_data := utils.ServiceSessionDataToSessionData(*svs)
	return session_data, nil
}
