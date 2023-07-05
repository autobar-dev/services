package utils

import (
	"os"
	"strings"

	"go.a5r.dev/services/realtime/types"
)

func LoadMeta() *types.Meta {
	commit_sha := ""
	version := ""

	if commit_sha_bytes, err := os.ReadFile(".meta/COMMIT_SHA"); err == nil {
		commit_sha = strings.TrimSpace(string(commit_sha_bytes))
	}

	if version_bytes, err := os.ReadFile(".meta/VERSION"); err == nil {
		version = strings.TrimSpace(string(version_bytes))
	}

	return &types.Meta{
		Hash:    commit_sha,
		Version: version,
	}
}
