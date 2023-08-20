package controllers

import (
	"errors"
	"time"

	"github.com/autobar-dev/services/user/types"
	"github.com/autobar-dev/services/user/utils"
	"github.com/autobar-dev/shared-libraries/go/auth-repository"
)

func Register(
	ac *types.AppContext,
	email string,
	password string,
	first_name string,
	last_name string,
	date_of_birth string,
	nationality string,
	locale string,
) error {
	urr := ac.Repositories.UnfinishedRegistration
	ar := authrepository.NewAuthRepository(ac.Config.AuthServiceURL, types.MICROSERVICE_NAME)
	etr := emailtemplaterepository.NewEmailTemplateRepository(ac.Config.EmailTemplateServiceURL, types.MICROSERVICE_NAME)
	er := emailrepository.NewEmailRepository(ac.Config.EmailServiceURL, types.MICROSERVICE_NAME)
	tr := translationrepository.NewTranslationRepository(ac.Config.TranslationServiceURL, types.MICROSERVICE_NAME)

	is_first_name_valid := utils.ValidateFirstName(first_name)
	if is_first_name_valid == false {
		return errors.New("invalid first name provided")
	}
	is_last_name_valid := utils.ValidateLastName(last_name)
	if is_last_name_valid == false {
		return errors.New("invalid last name provided")
	}
	is_date_of_birth_valid := utils.ValidateDateOfBirth(date_of_birth)
	if is_date_of_birth_valid == false {
		return errors.New("invalid date of birth provided. make sure the format follows yyyy-mm-dd")
	}
	is_locale_valid := utils.ValidateLocale(locale)
	if is_locale_valid == false {
		return errors.New("invalid locale provided")
	}

	// check both in user and unfinished registration repositories if e-mail has already been used
	is_used, err := IsEmailInUse(ac, email)
	if err != nil {
		return err
	}

	if !is_used {
		return errors.New("email is already in use")
	}

	// convert date of birth to time.Time and ignore error because it has already been validated
	dob, _ := utils.DateStringToTime(date_of_birth)

	// create auth-able user
	ar_autologin := false
	ar_remember_me := false

	_, err = ar.RegisterUser(email, password, &ar_autologin, &ar_remember_me)
	if err != nil {
		return err
	}

	// send confirmation e-mail
	confirmation_code := utils.RandomString(64, utils.LowercaseUppercaseNumbersSet)
	rendered_confirmation_email, err := etr.RenderEmailConfirmationTemplate(
		locale,
		first_name,
		last_name,
		confirmation_code,
		time.Now(),
	)

	// try to send email in background
	go func() {
		err := er.SendEmail(
			email,
			tr.GetTranslation("email_confirmation_subject", locale),
			rendered_confirmation_email,
		)
		if err != nil {
			ac.Logger.Error("failed to send email", err)
		}
	}()

	// create unfinished registration
	err = urr.Create(
		email,
		first_name,
		last_name,
		dob,
		locale,
		confirmation_code,
		time.Now(),
	)

	return err
}
