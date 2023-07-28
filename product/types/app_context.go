package types

import "github.com/autobar-dev/services/product/repositories"

type Repositories struct {
	Product     *repositories.ProductRepository
	SlugHistory *repositories.SlugHistoryRepository
	Cache       *repositories.CacheRepository
	Meili       *repositories.MeiliRepository
}

type AppContext struct {
	MetaFactors  *MetaFactors
	Config       *Config
	Repositories *Repositories
}
