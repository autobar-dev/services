package types

import (
	"github.com/autobar-dev/services/user/repositories"
	authrepository "github.com/autobar-dev/shared-libraries/go/auth-repository"
	emailrepository "github.com/autobar-dev/shared-libraries/go/email-repository"
	emailtemplaterepository "github.com/autobar-dev/shared-libraries/go/emailtemplate-repository"
	walletrepository "github.com/autobar-dev/shared-libraries/go/wallet-repository"
	"go.uber.org/zap"
)

type Repositories struct {
	User                   *repositories.UserRepository
	UnfinishedRegistration *repositories.UnfinishedRegistrationRepository
	Cache                  *repositories.CacheRepository
	Auth                   *authrepository.AuthRepository
	Email                  *emailrepository.EmailRepository
	EmailTemplate          *emailtemplaterepository.EmailTemplateRepository
	Wallet                 *walletrepository.WalletRepository
}

type AppContext struct {
	Logger       *zap.SugaredLogger
	MetaFactors  *MetaFactors
	Config       *Config
	Repositories *Repositories
}
