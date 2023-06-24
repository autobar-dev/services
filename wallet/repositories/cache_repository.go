package repositories

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisWallet struct {
	Id           int    `json:"id"`
	CurrencyCode string `json:"currency_code"`
	Balance      int    `json:"balance"`
}

type CacheRepository struct {
	redis *redis.Client
}

func NewCacheRepository(client *redis.Client) *CacheRepository {
	return &CacheRepository{
		redis: client,
	}
}

const walletCacheKeyPrefix = "wallet"

func generateWalletCacheKey(user_email string) string {
	return fmt.Sprintf("%s:%s", walletCacheKeyPrefix, user_email)
}

func (cr CacheRepository) GetWallet(user_email string) (*RedisWallet, error) {
	ctx := context.Background()

	rw_json, err := cr.redis.Get(ctx, generateWalletCacheKey(user_email)).Result()
	if err != nil {
		return nil, err
	}

	var rw RedisWallet
	_ = json.Unmarshal([]byte(rw_json), &rw)

	return &rw, nil
}

func (cr CacheRepository) SetWallet(id int, user_email string, balance int, currency_code string) error {
	rw := RedisWallet{
		Id:           id,
		CurrencyCode: currency_code,
		Balance:      balance,
	}
	rw_json_bytes, _ := json.Marshal(rw)
	rw_json := string(rw_json_bytes)

	ctx := context.Background()
	return cr.redis.Set(ctx, generateWalletCacheKey(user_email), rw_json, 0).Err()
}

func (cr CacheRepository) ClearWallet(user_email string) error {
	ctx := context.Background()
	return cr.redis.Del(ctx, generateWalletCacheKey(user_email)).Err()
}
