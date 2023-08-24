package types

import (
	"go.uber.org/zap"
)

type Repositories struct{}

type Providers struct {
	Auth AuthProvider
}

type AppContext struct {
	Logger       *zap.SugaredLogger
	MetaFactors  *MetaFactors
	Config       *Config
	Repositories *Repositories
	Providers    *Providers
}
