package types

import "time"

type MetaFactors struct {
	StartTime time.Time
	Hash      string
	Version   string
}

type Meta struct {
	Uptime  int64  `json:"uptime"`
	Hash    string `json:"hash"`
	Version string `json:"version"`
}
