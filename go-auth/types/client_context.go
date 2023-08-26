package types

type ClientContext struct {
	Type       TokenOwnerType `json:"sub_typ"`
	Identifier string         `json:"sub"`
	Role       *string        `json:"rol"`
}
