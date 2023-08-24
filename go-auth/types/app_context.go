package types

import (
	"github.com/autobar-dev/services/auth/repositories"
	"github.com/autobar-dev/shared-libraries/go/user-repository"
	"go.uber.org/zap"
)

type Repositories struct {
	RefreshToken *repositories.RefreshTokenRepository
	AuthUser     *repositories.AuthUserRepository
	AuthModule   *repositories.AuthModuleRepository
	User         *userrepository.UserRepository
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
