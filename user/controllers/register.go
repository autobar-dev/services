package controllers

import (
	"errors"
	"fmt"
	"time"

	"github.com/autobar-dev/services/user/types"
	"github.com/autobar-dev/services/user/utils"
	"github.com/google/uuid"
)

func Register(
	ac *types.AppContext,
	email string,
	password string,
	locale string,
) error {
	urr := ac.Repositories.UnfinishedRegistration
	ar := ac.Repositories.Auth
	etr := ac.Repositories.EmailTemplate
	er := ac.Repositories.Email

	// check both in user and unfinished registration repositories if e-mail has already been used
	is_used, err := IsEmailInUse(ac, email)
	if err != nil {
		return err
	}

	if is_used {
		return errors.New("email is already in use")
	}

	// create auth-able user
	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	user_id := uuid.String()

	err = ar.RegisterUser(user_id, email, password)
	if err != nil {
		return err
	}

	// send confirmation e-mail
	confirmation_code := utils.RandomString(64, utils.LowercaseUppercaseNumbersSet)
	email_template_params := map[string]interface{}{
		"email_confirmation_url": fmt.Sprintf("https://autobar.co/confirm-email?code=%s", confirmation_code),
	}

	rendered_confirmation_email, err := etr.RenderTemplate(
		"email-confirmation",  // Template name
		nil,                   // Template version (nil means latest)
		locale,                // Locale
		email_template_params, // Template params
	)
	if err != nil {
		return err
	}

	// try to send email in background
	go func() {
		err := er.Send(
			"noreply@autobar.co",
			email,
			"Autobar - Confirm your e-mail",
			rendered_confirmation_email.Plain,
			rendered_confirmation_email.Html,
		)
		if err != nil {
			ac.Logger.Error("failed to send email", err)
		}
	}()

	// create unfinished registration
	email_confirmation_expires_at := time.Now().Add(types.EmailConfirmationExpiryTime)
	err = urr.Create(
		user_id,
		email,
		locale,
		confirmation_code,
		email_confirmation_expires_at,
	)

	return err
}
