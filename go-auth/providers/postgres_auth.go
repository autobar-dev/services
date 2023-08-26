package providers

import (
	"errors"
	"fmt"
	"time"

	"github.com/autobar-dev/services/auth/repositories"
	"github.com/autobar-dev/services/auth/types"
	"github.com/autobar-dev/services/auth/utils"
	"golang.org/x/crypto/bcrypt"
)

const (
	RefreshTokenLength                      = 512
	UserRefreshTokenValidDuration           = time.Duration(time.Hour * 6)       // 6 hours
	UserRefreshTokenValidRememberMeDuration = time.Duration(time.Hour * 24 * 30) // 30 days
	ModuleRefreshTokenValidDuration         = time.Duration(time.Hour * 24 * 30) // 30 days
)

type PostgresAuthProvider struct {
	user_repository          *repositories.AuthUserRepository
	module_repository        *repositories.AuthModuleRepository
	refresh_token_repository *repositories.RefreshTokenRepository
	jwt_secret               string
}

func NewPostgresAuthProvider(
	ur *repositories.AuthUserRepository,
	mr *repositories.AuthModuleRepository,
	rtr *repositories.RefreshTokenRepository,
	jwt_secret string,
) *PostgresAuthProvider {
	return &PostgresAuthProvider{
		user_repository:          ur,
		module_repository:        mr,
		refresh_token_repository: rtr,
		jwt_secret:               jwt_secret,
	}
}

func (p *PostgresAuthProvider) LoginUser(
	email string,
	password string,
	remember_me bool,
) (refresh_token *string, err error) {
	fmt.Printf("getting user with email %s\n", email)

	auth_user, err := p.user_repository.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	fmt.Printf("validating password for user %s\n", auth_user.Id)

	err = bcrypt.CompareHashAndPassword([]byte(auth_user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	fmt.Printf("validation successful, generating refresh token for user %s\n", auth_user.Id)

	return p.createUserRefreshToken(auth_user.Id, remember_me)
}

func (p *PostgresAuthProvider) RegisterUser(
	user_id string,
	email string,
	password string,
) error {
	password_hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return p.user_repository.Create(user_id, email, string(password_hash))
}

func (p *PostgresAuthProvider) LoginModule(
	serial_number string,
	private_key string,
) (refresh_token *string, err error) {
	auth_module, err := p.module_repository.GetBySerialNumber(serial_number)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(auth_module.PrivateKey), []byte(private_key))
	if err != nil {
		return nil, errors.New("invalid private key")
	}

	return p.createModuleRefreshToken(serial_number)
}

func (p *PostgresAuthProvider) RegisterModule(
	serial_number string,
	private_key string,
) error {
	private_key_hash, err := bcrypt.GenerateFromPassword([]byte(private_key), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return p.module_repository.Create(serial_number, string(private_key_hash))
}

func (p *PostgresAuthProvider) InvalidateRefreshTokenById(
	token_id string,
) error {
	return p.refresh_token_repository.DeleteById(token_id)
}

func (p *PostgresAuthProvider) InvalidateRefreshTokenByToken(
	refresh_token string,
) error {
	return p.refresh_token_repository.DeleteByToken(refresh_token)
}

func (p *PostgresAuthProvider) GetRefreshTokenOwner(
	refresh_token string,
) (*types.RefreshTokenOwner, error) {
	rt, err := p.refresh_token_repository.GetByToken(refresh_token)
	if err != nil {
		return nil, err
	}

	if rt.ExpiresAt.Before(time.Now().UTC()) {
		return nil, errors.New("invalid refresh token")
	}

	var owner_type types.TokenOwnerType
	var identifier string

	if rt.UserId != nil {
		owner_type = types.UserTokenOwnerType
		identifier = *rt.UserId
	} else {
		owner_type = types.ModuleTokenOwnerType
		identifier = *rt.ModuleSerialNumber
	}

	return &types.RefreshTokenOwner{
		Type:       owner_type,
		Identifier: identifier,
	}, nil
}

func (p *PostgresAuthProvider) createUserRefreshToken(
	user_id string,
	remember_me bool,
) (refresh_token *string, err error) {
	var time_valid time.Duration
	if remember_me {
		time_valid = UserRefreshTokenValidRememberMeDuration
	} else {
		time_valid = UserRefreshTokenValidDuration
	}

	valid_until := time.Now().UTC().Add(time_valid)
	token := utils.RandomString(RefreshTokenLength, utils.RefreshTokenCharacters)

	err = p.refresh_token_repository.CreateForUser(user_id, token, valid_until)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (p *PostgresAuthProvider) createModuleRefreshToken(
	serial_number string,
) (refresh_token *string, err error) {
	valid_until := time.Now().UTC().Add(ModuleRefreshTokenValidDuration)
	token := utils.RandomString(RefreshTokenLength, utils.RefreshTokenCharacters)

	err = p.refresh_token_repository.CreateForModule(serial_number, token, valid_until)
	if err != nil {
		return nil, err
	}

	return &token, nil
}
