package repositories

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) *RedisRepository {
	return &RedisRepository{client: client}
}

func (rr RedisRepository) GetListenersCountForExchange(stream_name string) (*int, error) {
	ctx := context.Background()

	response, err := rr.client.Get(ctx, stream_name).Result()
	if err != nil {
		return nil, err
	}

	listeners_count, err := strconv.Atoi(response)
	if err != nil {
		return nil, err
	}

	return &listeners_count, nil
}

func (rr RedisRepository) IncrementListenersCountForExchange(stream_name string) error {
	ctx := context.Background()

	err := rr.client.Incr(ctx, stream_name).Err()
	if err != nil {
		return err
	}

	return nil
}

func (rr RedisRepository) DecrementListenersCountForExchange(stream_name string) error {
	ctx := context.Background()

	err := rr.client.Decr(ctx, stream_name).Err()
	if err != nil {
		return err
	}

	return nil
}
