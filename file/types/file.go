package types

import "time"

type File struct {
	Id        string    `json:"id"`
	Extension string    `json:"extension"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}
