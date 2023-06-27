package types

type ClientType string

const (
	ModuleClientType ClientType = "module"
	UserClientType   ClientType = "user"
)

type SessionData struct {
	ClientIdentifier string     `json:"client_identifier"`
	ClientType       ClientType `json:"client_type"`
}
