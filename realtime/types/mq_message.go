package types

type MqMessageType string

const (
	SimpleMessageType  MqMessageType = "simple"
	CommandMessageType MqMessageType = "command"
)
