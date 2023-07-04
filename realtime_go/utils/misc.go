package utils

import (
	"fmt"

	"go.a5r.dev/services/realtime/types"
)

func StreamNameFromClientInfo(ct types.ClientType, identifier string) string {
	return fmt.Sprintf("%s-%s", ct, identifier)
}
