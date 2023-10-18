package types

import "github.com/autobar-dev/services/file/repositories"

type Repositories struct {
	Cache *repositories.CacheRepository
	File  *repositories.FileRepository
	S3    *repositories.S3Repository
}

type AppContext struct {
	MetaFactors  *MetaFactors
	Config       *Config
	Repositories *Repositories
}
