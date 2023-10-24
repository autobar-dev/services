package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisProductBadgeType string

const (
	RedisProductBadgeTypePrimary   RedisProductBadgeType = "primary"
	RedisProductBadgeTypeSecondary RedisProductBadgeType = "secondary"
)

type RedisProductBadge struct {
	Type  RedisProductBadgeType `json:"type"`
	Label string                `json:"label"`
	Value *string               `json:"value"`
}

type RedisProduct struct {
	Id           string              `json:"id"`
	Names        map[string]string   `json:"names"`
	Descriptions map[string]string   `json:"descriptions"`
	CoverId      string              `json:"cover_id"`
	Enabled      bool                `json:"enabled"`
	Badges       []RedisProductBadge `json:"badges"`
	CreatedAt    time.Time           `json:"created_at"`
	UpdatedAt    time.Time           `json:"updated_at"`
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

const allProductsCacheKey = "product:all_products"

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

	pipe := cr.redis.Pipeline()
	for _, slug := range slugs {
		pipe.Del(ctx, generateSlugToLatestSlugCacheKey(slug))
	}
	_, err := pipe.Exec(ctx)

	return err
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

func (cr CacheRepository) SetProduct(
	id string,
	names map[string]string,
	descriptions map[string]string,
	cover_id string,
	enabled bool,
	badges []RedisProductBadge,
	created_at time.Time,
	updated_at time.Time,
) error {
	rp := RedisProduct{
		Id:           id,
		Names:        names,
		Descriptions: descriptions,
		CoverId:      cover_id,
		Enabled:      enabled,
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

func (cr CacheRepository) GetAllProducts() (*[]RedisProduct, error) {
	ctx := context.Background()

	rps_compressed, err := cr.redis.Get(ctx, allProductsCacheKey).Result()
	if err != nil {
		return nil, err
	}

	rps_json, err := DecompressBytes([]byte(rps_compressed))
	if err != nil {
		return nil, err
	}

	var rps []RedisProduct
	_ = json.Unmarshal(rps_json, &rps)

	return &rps, nil
}

func (cr CacheRepository) SetAllProducts(products []RedisProduct) error {
	rps_json_bytes, _ := json.Marshal(products)

	// Compress
	rps_compressed, err := CompressBytes(rps_json_bytes)
	if err != nil {
		return err
	}

	ctx := context.Background()
	return cr.redis.Set(ctx, allProductsCacheKey, rps_compressed, 0).Err()
}

func (cr CacheRepository) ClearAllProducts() error {
	ctx := context.Background()
	return cr.redis.Del(ctx, allProductsCacheKey).Err()
}
