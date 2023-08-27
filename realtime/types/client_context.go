package types

import authrepository "github.com/autobar-dev/shared-libraries/go/auth-repository"

type ClientContext struct {
	Type       authrepository.TokenOwnerType
	Identifier string
	Role       *string
}
