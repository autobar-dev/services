package types

type Message struct {
	Id      string `json:"id"`
	Command string `json:"command"`
	Args    string `json:"args"`
}

type Reply struct {
	Id string `json:"id"`
}
