package types

import (
	"github.com/go-oauth2/oauth2/v4/server"
	"go.uber.org/zap"
)

type Repositories struct {
	OAuthServer *server.Server
}

type AppContext struct {
	Logger       *zap.SugaredLogger
	MetaFactors  *MetaFactors
	Config       *Config
	Repositories *Repositories
}
