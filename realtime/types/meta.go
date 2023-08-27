package types

import "time"

type MetaFactors struct {
	Hash      string
	Version   string
	StartTime time.Time
}

type Meta struct {
	Uptime  int64  `json:"uptime"`
	Hash    string `json:"hash"`
	Version string `json:"version"`
}
