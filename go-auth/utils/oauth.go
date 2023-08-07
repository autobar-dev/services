package utils

import (
	"context"
	"time"

	"github.com/autobar-dev/services/auth/types"
	oauth_errors "github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/golang-jwt/jwt"
	"github.com/jackc/pgx/v4"
	oauth_pg "github.com/vgarvardt/go-oauth2-pg/v4"
	"github.com/vgarvardt/go-pg-adapter/pgx4adapter"
)

func SetupOAuthManager(ac *types.AppContext) (*manage.Manager, error) {
	ctx := context.Background()
	pgx_conn, err := pgx.Connect(ctx, ac.Config.DatabaseURL)

	adapter := pgx4adapter.NewConn(pgx_conn)

	token_store, err := oauth_pg.NewTokenStore(adapter, oauth_pg.WithTokenStoreGCInterval(time.Minute))
	if err != nil {
		return nil, err
	}

	client_store, err := oauth_pg.NewClientStore(adapter)
	if err != nil {
		return nil, err
	}

	manager := manage.NewDefaultManager()

	manager.MapTokenStorage(token_store)
	manager.MapClientStorage(client_store)

	manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte(ac.Config.JwtKey), jwt.SigningMethodHS512))

	return manager, nil
}

func SetupOAuthServer(ac *types.AppContext, manager *manage.Manager) (*server.Server, error) {
	server := server.NewDefaultServer(manager)

	server.SetAllowGetAccessRequest(true)
	server.SetInternalErrorHandler(func(err error) (re *oauth_errors.Response) {
		ac.Logger.Warnf("OAuth internal server error", err)
		return
	})

	return server, nil
}
