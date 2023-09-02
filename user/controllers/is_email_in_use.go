package controllers

import "github.com/autobar-dev/services/user/types"

func IsEmailInUse(ac *types.AppContext, email string) (bool, error) {
	ur := ac.Repositories.User
	urr := ac.Repositories.UnfinishedRegistration

	_, err := ur.GetByEmail(email)
	if err != nil { // if err != nil, email is not in use in the user table
		_, err = urr.GetByEmail(email)
		if err != nil { // if err != nil, email is not in use in the unfinished_registrations table
			return false, nil
		}
	}
	return true, nil
}
