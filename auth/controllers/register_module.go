package controllers

import (
	"github.com/autobar-dev/services/auth/types"
	"github.com/autobar-dev/services/auth/utils"
)

func RegisterModule(ac *types.AppContext, serial_number string) (private_key *string, err error) {
	ap := ac.Providers.Auth

	generated_private_key := utils.RandomString(72, utils.PrivateKeyCharacters) // max 72 bytes because of bcrypt

	err = ap.RegisterModule(serial_number, generated_private_key)
	if err != nil {
		return nil, err
	}

	return &generated_private_key, nil
}
