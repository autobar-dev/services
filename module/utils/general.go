package utils

import (
	"github.com/autobar-dev/services/module/types"
	authrepository "github.com/autobar-dev/shared-libraries/go/auth-repository"
)

func ServiceClientTypeToClientType(sct authrepository.TokenOwnerType) *types.ClientType {
	var ct types.ClientType

	switch sct {
	case authrepository.ModuleTokenOwnerType:
		ct = types.ModuleClientType
	case authrepository.UserTokenOwnerType:
		ct = types.UserClientType
	}

	return &ct
}
