package utils

import (
	"errors"
	"fmt"

	"go.a5r.dev/services/realtime/types"
)

func ExchangeNameFromClientInfo(ct types.ClientType, identifier string) string {
	return fmt.Sprintf("%s-%s", ct, identifier)
}

func ClientTypeFromString(cts string) (*types.ClientType, error) {
	var ct types.ClientType
	switch cts {
	case "module":
		ct = types.ModuleClientType
		break
	case "user":
		ct = types.UserClientType
		break
	default:
		return nil, errors.New("failed to parse client type")
	}

	return &ct, nil
}
