package utils

import (
	"os"
	"strings"
	"time"

	"github.com/autobar-dev/services/realtime/types"
)

func GetMetaFromFactors(mf *types.MetaFactors) *types.Meta {
	now := time.Now()
	uptime := int64(now.Sub(mf.StartTime).Seconds())

	return &types.Meta{
		Uptime:  uptime,
		Hash:    mf.Hash,
		Version: mf.Version,
	}
}

func GetMetaFactors() *types.MetaFactors {
	commit_sha := ""
	version := ""

	if commit_sha_bytes, err := os.ReadFile(".meta/HASH"); err == nil {
		commit_sha = strings.TrimSpace(string(commit_sha_bytes))
	}

	if version_bytes, err := os.ReadFile(".meta/VERSION"); err == nil {
		version = strings.TrimSpace(string(version_bytes))
	}

	start_time := time.Now()

	return &types.MetaFactors{
		StartTime: start_time,
		Hash:      commit_sha,
		Version:   version,
	}
}
