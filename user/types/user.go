package types

import "time"

type User struct {
	Id                         string    `json:"id"`
	Email                      string    `json:"email"`
	FirstName                  string    `json:"first_name"`
	LastName                   string    `json:"last_name"`
	DateOfBirth                time.Time `json:"date_of_birth"`
	Locale                     string    `json:"locale"`
	IdentityVerificationId     string    `json:"identity_verification_id"`
	IdentityVerificationSource string    `json:"identity_verification_source"`
	CreatedAt                  time.Time `json:"created_at"`
	UpdatedAt                  time.Time `json:"updated_at"`
}
