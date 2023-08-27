package utils

import (
	"errors"
	"fmt"

	"github.com/autobar-dev/services/realtime/types"
)

func ExchangeNameFromClientInfo(ct types.ClientType, identifier string) string {
	return fmt.Sprintf("%s-%s", ct, identifier)
}

func ReplyExchangeNameFromClientInfo(ct types.ClientType, identifier string) string {
	return fmt.Sprintf("replies-%s-%s", ct, identifier)
}

func MessageConsumerName(message_id string) string {
	return fmt.Sprintf("msg_consumer-%s", message_id)
}

func QueueConsumerName(queue_name string) string {
	return fmt.Sprintf("queue_consumer-%s", queue_name)
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
