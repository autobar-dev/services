package utils

import (
	"github.com/autobar-dev/services/realtime/types"
	"github.com/autobar-dev/shared-libraries/go/auth-repository"
)

func ServiceClientTypeToClientType(sct authrepository.TokenOwnerType) types.ClientType {
	var ct types.ClientType

	switch sct {
	case authrepository.UserTokenOwnerType:
		ct = types.UserClientType
		break
	case authrepository.ModuleTokenOwnerType:
		ct = types.ModuleClientType
		break
	}

	return ct
}
