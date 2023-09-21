package repositories

import (
	"context"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type StateRepository struct {
	client *redis.Client
}

func NewStateRepository(client *redis.Client) *StateRepository {
	return &StateRepository{client: client}
}

func generateListenersCountForExchangeKey(stream_name string) string {
	return fmt.Sprintf("realtime:listeners_count:%s", stream_name)
}

func (rr StateRepository) GetListenersCountForExchange(stream_name string) (*int, error) {
	ctx := context.Background()

	response, err := rr.client.Get(ctx, generateListenersCountForExchangeKey(stream_name)).Result()
	if err != nil {
		return nil, err
	}

	listeners_count, err := strconv.Atoi(response)
	if err != nil {
		return nil, err
	}

	return &listeners_count, nil
}

func (rr StateRepository) IncrementListenersCountForExchange(stream_name string) error {
	ctx := context.Background()

	err := rr.client.Incr(ctx, generateListenersCountForExchangeKey(stream_name)).Err()
	if err != nil {
		return err
	}

	return nil
}

func (rr StateRepository) DecrementListenersCountForExchange(stream_name string) error {
	ctx := context.Background()

	err := rr.client.Decr(ctx, generateListenersCountForExchangeKey(stream_name)).Err()
	if err != nil {
		return err
	}

	return nil
}
