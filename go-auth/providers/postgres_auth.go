package providers

import (
	"github.com/autobar-dev/services/auth/types"
	"github.com/jmoiron/sqlx"
)

type PostgresAuthProvider struct {
	db *sqlx.DB
}

func NewPostgresAuthProvider(db *sqlx.DB) *PostgresAuthProvider {
	return &PostgresAuthProvider{db}
}

func (p *PostgresAuthProvider) Login(email string, password string, remember_me bool) (*types.Tokens, error) {
	return nil, nil
}

func (p *PostgresAuthProvider) Logout(refresh_token string) error {
	return nil
}

func (p *PostgresAuthProvider) Refresh(refresh_token string) (*types.Tokens, error) {
	return nil, nil
}

func (p *PostgresAuthProvider) Register(email string, password string) (*types.Tokens, error) {
	return nil, nil
}
