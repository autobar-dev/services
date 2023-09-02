package types

import (
	"time"

	walletrepository "github.com/autobar-dev/shared-libraries/go/wallet-repository"
)

type UserExtended struct {
	Id                         string                  `json:"id"`
	Email                      string                  `json:"email"`
	FirstName                  string                  `json:"first_name"`
	LastName                   string                  `json:"last_name"`
	DateOfBirth                time.Time               `json:"date_of_birth"`
	Locale                     string                  `json:"locale"`
	Wallet                     walletrepository.Wallet `json:"wallet"`
	IdentityVerificationId     *string                 `json:"identity_verification_id"`
	IdentityVerificationSource *string                 `json:"identity_verification_source"`
	CreatedAt                  time.Time               `json:"created_at"`
	UpdatedAt                  time.Time               `json:"updated_at"`
}

type User struct {
	Id                         string    `json:"id"`
	Email                      string    `json:"email"`
	FirstName                  string    `json:"first_name"`
	LastName                   string    `json:"last_name"`
	DateOfBirth                time.Time `json:"date_of_birth"`
	Locale                     string    `json:"locale"`
	IdentityVerificationId     *string   `json:"identity_verification_id"`
	IdentityVerificationSource *string   `json:"identity_verification_source"`
	CreatedAt                  time.Time `json:"created_at"`
	UpdatedAt                  time.Time `json:"updated_at"`
}
