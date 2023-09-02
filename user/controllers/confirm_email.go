package controllers

import (
	"time"

	"github.com/autobar-dev/services/user/types"
)

func ConfirmEmail(
	ac *types.AppContext,
	confirmation_code string,
	first_name string,
	last_name string,
	date_of_birth time.Time,
	currency_code string,
) error {
	urr := ac.Repositories.UnfinishedRegistration
	ur := ac.Repositories.User
	wr := ac.Repositories.Wallet

	err := IsConfirmationCodeValid(ac, confirmation_code)
	if err != nil {
		return err
	}

	unfinished_registration, err := urr.GetByConfirmationCode(confirmation_code)
	if err != nil {
		return err
	}

	// create user
	err = ur.Create(
		unfinished_registration.Id,
		unfinished_registration.Email,
		first_name,
		last_name,
		date_of_birth,
		unfinished_registration.Locale,
	)
	if err != nil {
		return err
	}

	// create wallet
	_, err = wr.Create(
		unfinished_registration.Id,
		currency_code,
	)
	if err != nil {
		return err
	}

	// delete unfinished registration
	go func() {
		err := urr.Delete(unfinished_registration.Id)
		if err != nil {
			ac.Logger.Error("failed to delete unfinished registration", err)
		}
	}()

	return nil
}
