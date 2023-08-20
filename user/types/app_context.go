package types

import (
	"github.com/autobar-dev/services/user/repositories"
	"go.uber.org/zap"
)

type Repositories struct {
	User                   *repositories.UserRepository
	UnfinishedRegistration *repositories.UnfinishedRegistrationRepository
	Cache                  *repositories.CacheRepository
}

type AppContext struct {
	Logger       *zap.SugaredLogger
	MetaFactors  *MetaFactors
	Config       *Config
	Repositories *Repositories
}
