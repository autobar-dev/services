package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisActivationSession struct {
	Id                string    `json:"id"`
	UserId            string    `json:"user_id"`
	SerialNumber      string    `json:"serial_number"`
	ProductId         string    `json:"product_id"`
	Price             int       `json:"price"`
	AmountMillilitres int       `json:"amount_millilitres"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type StateRepository struct {
	redis *redis.Client
}

func NewStateRepository(client *redis.Client) *StateRepository {
	return &StateRepository{
		redis: client,
	}
}

// OTK for module
const otkForModuleCacheKey = "module:otk"

func generateOtkForModuleCacheKey(serial_number string) string {
	return fmt.Sprintf("%s:%s", otkForModuleCacheKey, serial_number)
}

func (sr StateRepository) GetOtkForModule(serial_number string) (*string, error) {
	ctx := context.Background()

	otk, err := sr.redis.Get(ctx, generateOtkForModuleCacheKey(serial_number)).Result()
	if err != nil {
		return nil, err
	}

	return &otk, nil
}

func (sr StateRepository) SetOtkForModule(
	serial_number string,
	otk string,
) error {
	ctx := context.Background()
	return sr.redis.Set(ctx, generateOtkForModuleCacheKey(serial_number), otk, 0).Err()
}

func (sr StateRepository) ClearOtkForModule(serial_number string) error {
	ctx := context.Background()
	return sr.redis.Del(ctx, generateOtkForModuleCacheKey(serial_number)).Err()
}

// Activation session
const activationSessionCacheKey = "module:activation_session"

func generateActivationSessionCacheKey(id string) string {
	return fmt.Sprintf("%s:%s", activationSessionCacheKey, id)
}

func (sr StateRepository) GetActivationSession(id string) (*RedisActivationSession, error) {
	ctx := context.Background()

	ras_json, err := sr.redis.Get(ctx, generateActivationSessionCacheKey(id)).Result()
	if err != nil {
		return nil, err
	}

	var ras RedisActivationSession
	_ = json.Unmarshal([]byte(ras_json), &ras)

	return &ras, nil
}

func (sr StateRepository) SetActivationSession(
	id string,
	activation_session *RedisActivationSession,
) error {
	ras_json_bytes, _ := json.Marshal(activation_session)

	ctx := context.Background()
	return sr.redis.Set(ctx, generateActivationSessionCacheKey(id), ras_json_bytes, 0).Err()
}

func (sr StateRepository) ClearActivationSession(id string) error {
	ctx := context.Background()
	return sr.redis.Del(ctx, generateActivationSessionCacheKey(id)).Err()
}

// Activation session ID for module
const activationSessionIdForModuleCacheKey = "module:activation_session_id_for_module"

func generateActivationSessionIdForModuleCacheKey(serial_number string) string {
	return fmt.Sprintf("%s:%s", activationSessionIdForModuleCacheKey, serial_number)
}

func (sr StateRepository) GetActivationSessionIdForModule(serial_number string) (*string, error) {
	ctx := context.Background()

	id, err := sr.redis.Get(ctx, generateActivationSessionIdForModuleCacheKey(serial_number)).Result()
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (sr StateRepository) SetActivationSessionIdForModule(
	serial_number string,
	id string,
) error {
	ctx := context.Background()
	return sr.redis.Set(ctx, generateActivationSessionIdForModuleCacheKey(serial_number), id, 0).Err()
}

func (sr StateRepository) ClearActivationSessionIdForModule(serial_number string) error {
	ctx := context.Background()
	return sr.redis.Del(ctx, generateActivationSessionIdForModuleCacheKey(serial_number)).Err()
}

// Activation session ID for user
const activationSessionIdForUserCacheKey = "module:activation_session_id_for_user"

func generateActivationSessionIdForUserCacheKey(user_id string) string {
	return fmt.Sprintf("%s:%s", activationSessionIdForUserCacheKey, user_id)
}

func (sr StateRepository) GetActivationSessionIdForUser(user_id string) (*string, error) {
	ctx := context.Background()

	id, err := sr.redis.Get(ctx, generateActivationSessionIdForUserCacheKey(user_id)).Result()
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (sr StateRepository) SetActivationSessionIdForUser(
	user_id string,
	id string,
) error {
	ctx := context.Background()
	return sr.redis.Set(ctx, generateActivationSessionIdForUserCacheKey(user_id), id, 0).Err()
}

func (sr StateRepository) ClearActivationSessionIdForUser(user_id string) error {
	ctx := context.Background()
	return sr.redis.Del(ctx, generateActivationSessionIdForUserCacheKey(user_id)).Err()
}
