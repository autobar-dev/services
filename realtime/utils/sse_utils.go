package utils

import (
	"encoding/json"
	"time"

	"github.com/autobar-dev/services/realtime/types"
	"github.com/r3labs/sse/v2"
)

func CreateCommandSseEvent(id string, command_name string, args string) *sse.Event {
	command := &types.Message{
		Id:      id,
		Command: command_name,
		Args:    args,
	}

	command_json, _ := json.Marshal(command)

	event_name := []byte("command")
	data := command_json

	return &sse.Event{
		Event: event_name,
		Data:  data,
	}
}

func CreateHeartbeatSseEvent() *sse.Event {
	now := time.Now().UTC()
	now_string := now.Format(time.RFC3339Nano)

	event_name := []byte("heartbeat")
	data := []byte(now_string)

	return &sse.Event{
		Event: event_name,
		Data:  data,
	}
}
