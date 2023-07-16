package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisProduct struct {
	Id           string            `json:"id"`
	Names        map[string]string `json:"names"`
	Descriptions map[string]string `json:"descriptions"`
	Cover        *string           `json:"cover"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
}

type CacheRepository struct {
	redis *redis.Client
}

func NewCacheRepository(client *redis.Client) *CacheRepository {
	return &CacheRepository{
		redis: client,
	}
}

const latestSlugToProductIdCacheKey = "product:slug_product_id"

func generateLatestSlugToProductIdCacheKey(ls string) string {
	return fmt.Sprintf("%s:%s", latestSlugToProductIdCacheKey, ls)
}

const slugToLatestSlugCacheKey = "product:latest_slug"

func generateSlugToLatestSlugCacheKey(slug string) string {
	return fmt.Sprintf("%s:%s", slugToLatestSlugCacheKey, slug)
}

const productIdToProductCacheKey = "product"

func generateProductIdToProductCacheKey(id string) string {
	return fmt.Sprintf("%s:%s", productIdToProductCacheKey, id)
}

func (cr CacheRepository) GetProductIdFromLatestSlug(ls string) (*string, error) {
	ctx := context.Background()

	pid, err := cr.redis.Get(ctx, generateLatestSlugToProductIdCacheKey(ls)).Result()
	if err != nil {
		return nil, err
	}

	return &pid, nil
}

func (cr CacheRepository) SetProductIdFromLatestSlug(ls string, pid string) error {
	ctx := context.Background()
	return cr.redis.Set(ctx, generateLatestSlugToProductIdCacheKey(ls), pid, 0).Err()
}

func (cr CacheRepository) ClearProductIdFromLatestSlug(ls string) error {
	ctx := context.Background()
	return cr.redis.Del(ctx, generateLatestSlugToProductIdCacheKey(ls)).Err()
}

func (cr CacheRepository) GetLatestSlugFromSlug(slug string) (*string, error) {
	ctx := context.Background()

	ls, err := cr.redis.Get(ctx, generateSlugToLatestSlugCacheKey(slug)).Result()
	if err != nil {
		return nil, err
	}

	return &ls, nil
}

func (cr CacheRepository) SetMultipleSlugsToLatestSlug(all_slugs []string) error {
	ctx := context.Background()

	latest_slug := all_slugs[len(all_slugs)-1]

	mset_args := []string{}
	for _, slug := range all_slugs {
		mset_args = append(mset_args, generateSlugToLatestSlugCacheKey(slug))
		mset_args = append(mset_args, latest_slug)
	}

	return cr.redis.MSet(ctx, mset_args).Err()
}

func (cr CacheRepository) ClearMultipleSlugsToLatestSlug(slugs []string) error {
	ctx := context.Background()

	command := []string{"DEL"}
	del_args := []string{}
	for _, slug := range slugs {
		del_args = append(del_args, generateSlugToLatestSlugCacheKey(slug))
	}

	command = append(command, del_args...)

	return cr.redis.Do(ctx, command).Err()
}

func (cr CacheRepository) GetProduct(id string) (*RedisProduct, error) {
	ctx := context.Background()

	rp_compressed, err := cr.redis.Get(ctx, generateProductIdToProductCacheKey(id)).Result()
	if err != nil {
		return nil, err
	}

	rp_json, err := DecompressBytes([]byte(rp_compressed))
	if err != nil {
		return nil, err
	}

	var rp RedisProduct
	_ = json.Unmarshal(rp_json, &rp)

	return &rp, nil
}

func (cr CacheRepository) SetProduct(id string, names map[string]string, descriptions map[string]string, cover *string, created_at time.Time, updated_at time.Time) error {
	rp := RedisProduct{
		Id:           id,
		Names:        names,
		Descriptions: descriptions,
		Cover:        cover,
		CreatedAt:    created_at,
		UpdatedAt:    updated_at,
	}
	rp_json_bytes, _ := json.Marshal(rp)

	// Compress
	rp_compressed, err := CompressBytes(rp_json_bytes)
	if err != nil {
		return err
	}

	ctx := context.Background()
	return cr.redis.Set(ctx, generateProductIdToProductCacheKey(id), rp_compressed, 0).Err()
}

func (cr CacheRepository) ClearProduct(id string) error {
	ctx := context.Background()

	return cr.redis.Del(ctx, generateProductIdToProductCacheKey(id)).Err()
}
