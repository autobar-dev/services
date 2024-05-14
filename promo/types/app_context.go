package types

import "github.com/autobar-dev/services/promo/repositories"

type Repositories struct {
	Cache *repositories.CacheRepository
}

type AppContext struct {
	MetaFactors  *MetaFactors
	Config       *Config
	Repositories *Repositories
}
