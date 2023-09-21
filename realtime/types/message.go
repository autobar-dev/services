package types

type Command struct {
	Id      string                 `json:"id"`
	Command string                 `json:"command"`
	Args    map[string]interface{} `json:"args"`
}

type Reply struct {
	Id string `json:"id"`
}
