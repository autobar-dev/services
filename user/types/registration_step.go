package types

type RegistrationStep string

const (
  EMAIL_FIRST_NAME_LAST_NAME_LOCALE RegistrationStep = "email_first-name_last-name_locale"
  PHONE_NUMBER RegistrationStep = "phone-number"
  DATE_OF_BIRTH RegistrationStep = "date-of-birth"
)
