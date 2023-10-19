package types

import (
	"github.com/autobar-dev/services/product/repositories"
	filerepository "github.com/autobar-dev/shared-libraries/go/file-repository"
)

type Repositories struct {
	Product     *repositories.ProductRepository
	SlugHistory *repositories.SlugHistoryRepository
	Cache       *repositories.CacheRepository
	Meili       *repositories.MeiliRepository
	File        *filerepository.FileRepository
}

type AppContext struct {
	MetaFactors  *MetaFactors
	Config       *Config
	Repositories *Repositories
}
