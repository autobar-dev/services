package utils

import (
	"encoding/json"
	"time"

	"github.com/r3labs/sse/v2"
	"go.a5r.dev/services/realtime/types"
)

func CreateCommandSseEvent(command_name string, args string) *sse.Event {
	command := &types.Command{
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

func CreateSimpleSseEvent(body string) *sse.Event {
	event_name := []byte("simple")
	data := []byte(body)

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
