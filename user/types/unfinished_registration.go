package types

import "time"

type UnfinishedRegistration struct {
	Id                            string    `json:"id"`
	Email                         string    `json:"email"`
	FirstName                     string    `json:"first_name"`
	LastName                      string    `json:"last_name"`
	DateOfBirth                   string    `json:"date_of_birth"`
	Locale                        string    `json:"locale"`
	EmailConfirmed                bool      `json:"email_confirmed"`
	EmailConfirmationCode         string    `json:"email_confirmation_code"`
	EmailConfirmationCodeIssuedAt string    `json:"email_confirmation_code_issued_at"`
	UpdatedAt                     time.Time `json:"updated_at"`
	CreatedAt                     time.Time `json:"created_at"`
}
