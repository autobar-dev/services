package controllers

import (
	"errors"
	"time"

	"github.com/autobar-dev/services/user/types"
)

func IsConfirmationCodeValid(ac *types.AppContext, confirmation_code string) error {
	urr := ac.Repositories.UnfinishedRegistration

	unfinished_registration, err := urr.GetByConfirmationCode(confirmation_code)
	if err != nil {
		return errors.New("confirmation code is invalid")
	}

	if time.Now().UTC().After(unfinished_registration.EmailConfirmationCodeExpiresAt) {
		return errors.New("confirmation code has expired")
	}

	return nil
}
