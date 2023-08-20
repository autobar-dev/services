package controllers

import "github.com/autobar-dev/services/user/types"

func IsEmailInUse(ac *types.AppContext, email string) (bool, error) {
  ur := ac.Repositories.User
  urr := ac.Repositories.UnfinishedRegistration

  _, err := ur.GetByEmail(email)
  if err != nil {
    _, err = urr.GetByEmail(email)
    if err != nil {
      return false, nil
    }
  }
  return true, nil
}
