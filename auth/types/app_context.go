package types

import (
	"github.com/autobar-dev/services/auth/repositories"
	"go.uber.org/zap"
)

type Repositories struct {
	RefreshToken *repositories.RefreshTokenRepository
	AuthUser     *repositories.AuthUserRepository
	AuthModule   *repositories.AuthModuleRepository
}

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
