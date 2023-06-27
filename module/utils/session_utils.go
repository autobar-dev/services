package utils

import (
	"go.a5r.dev/services/module/repositories"
	"go.a5r.dev/services/module/types"
)

func ServiceClientTypeToClientType(sct repositories.ServiceAuthClientType) *types.ClientType {
	var ct types.ClientType

	switch sct {
	case repositories.ModuleServiceAuthClientType:
		ct = types.ModuleClientType
	case repositories.UserServiceAuthClientType:
		ct = types.UserClientType
	}

	return &ct
}

func ServiceSessionDataToSessionData(ssd repositories.ServiceSessionData) *types.SessionData {

	return &types.SessionData{
		ClientIdentifier: ssd.ClientIdentifier,
		ClientType:       *ServiceClientTypeToClientType(ssd.ClientType),
	}
}
