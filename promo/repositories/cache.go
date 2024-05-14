package repositories

import (
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisFile struct {
	Id        string    `json:"id"`
	Extension string    `json:"extension"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

type CacheRepository struct {
	redis *redis.Client
}

func NewCacheRepository(client *redis.Client) *CacheRepository {
	return &CacheRepository{
		redis: client,
	}
}
