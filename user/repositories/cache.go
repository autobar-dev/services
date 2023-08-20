package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisUser struct {
	Id                     string     `json:"id"`
	Email                  string     `json:"email"`
	PhoneNumberCountryCode string     `json:"phone_number_country_code"`
	PhoneNumber            string     `json:"phone_number"`
	FirstName              string     `json:"first_name"`
	LastName               string     `json:"last_name"`
	Locale                 string     `json:"locale"`
	Verified               bool       `json:"verified"`
	VerifiedAt             *time.Time `json:"verified_at"`
	CreatedAt              time.Time  `json:"created_at"`
	UpdatedAt              time.Time  `json:"updated_at"`
}

type CacheRepository struct {
	redis *redis.Client
}

func NewCacheRepository(client *redis.Client) *CacheRepository {
	return &CacheRepository{
		redis: client,
	}
}

const userByIdCacheKey = "user:by_id"

func generateUserByIdCacheKey(id string) string {
	return fmt.Sprintf("%s:%s", userByIdCacheKey, id)
}

func (cr CacheRepository) GetUser(id string) (*RedisUser, error) {
	ctx := context.Background()

	ru_compressed, err := cr.redis.Get(ctx, generateUserByIdCacheKey(id)).Result()
	if err != nil {
		return nil, err
	}

	rp_json, err := DecompressBytes([]byte(ru_compressed))
	if err != nil {
		return nil, err
	}

	var ru RedisUser
	_ = json.Unmarshal(rp_json, &ru)

	return &ru, nil
}

func (cr CacheRepository) SetUser(
	id string,
	email string,
	phone_number_country_code string,
	phone_number string,
	first_name string,
	last_name string,
	locale string,
	verified bool,
	verified_at *time.Time,
	created_at time.Time,
	updated_at time.Time,
) error {
	ru := RedisUser{
		Id:           id,
		Email: email,
		PhoneNumberCountryCode: phone_number_country_code,
		PhoneNumber: phone_number,
		FirstName: first_name,
		LastName: last_name,
		Locale: locale,
		Verified: verified,
		VerifiedAt: verified_at,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
	}
	ru_json_bytes, _ := json.Marshal(ru)

	// Compress
	ru_compressed, err := CompressBytes(ru_json_bytes)
	if err != nil {
		return err
	}

	ctx := context.Background()
	return cr.redis.Set(ctx, generateUserByIdCacheKey(id), ru_compressed, 0).Err()
}

func (cr CacheRepository) ClearUser(id string) error {
	ctx := context.Background()
	return cr.redis.Del(ctx, generateUserByIdCacheKey(id)).Err()
}

