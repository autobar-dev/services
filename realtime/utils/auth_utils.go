package utils

import (
	"go.a5r.dev/services/realtime/repositories"
	"go.a5r.dev/services/realtime/types"
)

func ServiceClientTypeToClientType(sct repositories.ServiceAuthClientType) types.ClientType {
	var ct types.ClientType

	switch sct {
	case repositories.UserServiceAuthClientType:
		ct = types.UserClientType
		break
	case repositories.ModuleServiceAuthClientType:
		ct = types.ModuleClientType
		break
	}

	return ct
}
