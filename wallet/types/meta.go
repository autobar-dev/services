package types

import "time"

type MetaFactors struct {
	Hash      string
	Version   string
	StartTime time.Time
}

type Meta struct {
	Hash    string `json:"hash"`
	Version string `json:"version"`
	Uptime  int64  `json:"uptime"`
}
