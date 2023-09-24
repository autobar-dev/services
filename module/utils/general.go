package utils

import (
	"encoding/json"

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

func StructToJsonMap(v any) (map[string]interface{}, error) {
	v_json, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	var parsed map[string]interface{}
	err = json.Unmarshal(v_json, &parsed)
	if err != nil {
		return nil, err
	}

	return parsed, nil
}
