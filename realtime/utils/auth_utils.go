package utils

import (
	"github.com/autobar-dev/shared-libraries/go/auth-repository"
	"go.a5r.dev/services/realtime/types"
)

func ServiceClientTypeToClientType(sct authrepository.ServiceAuthClientType) types.ClientType {
	var ct types.ClientType

	switch sct {
	case authrepository.UserServiceAuthClientType:
		ct = types.UserClientType
		break
	case authrepository.ModuleServiceAuthClientType:
		ct = types.ModuleClientType
		break
	}

	return ct
}
