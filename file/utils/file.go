package utils

import (
	"strings"

	"github.com/autobar-dev/services/file/repositories"
	"github.com/autobar-dev/services/file/types"
)

func FileExtensionFromFileName(filename string) string {
	return filename[strings.LastIndex(filename, ".")+1:]
}

func PostgresFileToFile(pf repositories.PostgresFile, url string) *types.File {
	return &types.File{
		Id:        pf.Id,
		Extension: pf.Extension,
		Url:       url,
		CreatedAt: pf.CreatedAt,
	}
}

func RedisFileToFile(rf repositories.RedisFile) *types.File {
	return &types.File{
		Id:        rf.Id,
		Extension: rf.Extension,
		Url:       rf.Url,
		CreatedAt: rf.CreatedAt,
	}
}

func FileToRedisFile(p types.File) *repositories.RedisFile {
	return &repositories.RedisFile{
		Id:        p.Id,
		Extension: p.Extension,
		Url:       p.Url,
		CreatedAt: p.CreatedAt,
	}
}
