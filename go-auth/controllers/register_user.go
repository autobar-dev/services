package controllers

import (
	"github.com/autobar-dev/services/auth/types"
)

func RegisterUser(ac *types.AppContext, user_id string, email string, password string) error {
	ap := ac.Providers.Auth

	return ap.RegisterUser(user_id, email, password)
}
