package types

import authrepository "github.com/autobar-dev/shared-libraries/go/auth-repository"

type ClientContext struct {
	Type       authrepository.TokenOwnerType `json:"sub_typ"`
	Identifier string                        `json:"sub"`
	Role       *string                       `json:"rol"`
}
