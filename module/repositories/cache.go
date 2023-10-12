package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	currencyrepository "github.com/autobar-dev/shared-libraries/go/currency-repository"
	"github.com/redis/go-redis/v9"
)

type RedisDisplayUnit struct {
	Id                     int32     `json:"id"`
	Amount                 float64   `json:"amount"`
	Symbol                 string    `json:"symbol"`
	DivisorFromMillilitres float64   `json:"divisor_from_millilitres"`
	DecimalsDisplayed      int32     `json:"decimals_displayed"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}

type RedisModule struct {
	Id              int32                       `json:"id"`
	SerialNumber    string                      `json:"serial_number"`
	StationId       *string                     `json:"station_id"`
	ProductId       *string                     `json:"product_id"`
	Enabled         bool                        `json:"enabled"`
	Prices          map[string]int              `json:"prices"`
	DisplayCurrency currencyrepository.Currency `json:"display_currency"`
	DisplayUnit     RedisDisplayUnit            `json:"display_unit"`
	CreatedAt       time.Time                   `json:"created_at"`
	UpdatedAt       time.Time                   `json:"updated_at"`
}

type CacheRepository struct {
	redis *redis.Client
}

func NewCacheRepository(client *redis.Client) *CacheRepository {
	return &CacheRepository{
		redis: client,
	}
}

const moduleCacheKey = "module"

func generateModuleCacheKey(serial_number string) string {
	return fmt.Sprintf("%s:%s", moduleCacheKey, serial_number)
}

const allModulesForStationCacheKey = "module:all_modules_for_station"

func generateAllModulesForStationCacheKey(station_id string) string {
	return fmt.Sprintf("%s:%s", allModulesForStationCacheKey, station_id)
}

func (cr *CacheRepository) GetModule(serial_number string) (*RedisModule, error) {
	ctx := context.Background()

	rm_compressed, err := cr.redis.Get(ctx, generateModuleCacheKey(serial_number)).Result()
	if err != nil {
		return nil, err
	}

	rm_json, err := DecompressBytes([]byte(rm_compressed))
	if err != nil {
		return nil, err
	}

	var rm RedisModule
	_ = json.Unmarshal(rm_json, &rm)

	return &rm, nil
}

func (cr *CacheRepository) SetModule(
	rm RedisModule,
) error {
	rm_json_bytes, _ := json.Marshal(rm)

	// Compress
	rm_compressed, err := CompressBytes(rm_json_bytes)
	if err != nil {
		return err
	}

	ctx := context.Background()
	return cr.redis.Set(ctx, generateModuleCacheKey(rm.SerialNumber), rm_compressed, 0).Err()
}

func (cr *CacheRepository) ClearModule(serial_number string) error {
	ctx := context.Background()
	return cr.redis.Del(ctx, generateModuleCacheKey(serial_number)).Err()
}

func (cr *CacheRepository) GetAllModulesForStation(station_id string) (*[]RedisModule, error) {
	ctx := context.Background()

	rms_compressed, err := cr.redis.Get(ctx, generateAllModulesForStationCacheKey(station_id)).Result()
	if err != nil {
		return nil, err
	}

	rms_json, err := DecompressBytes([]byte(rms_compressed))
	if err != nil {
		return nil, err
	}

	var rms []RedisModule
	_ = json.Unmarshal(rms_json, &rms)

	return &rms, nil
}

func (cr *CacheRepository) SetAllModulesForStation(station_id string, modules []RedisModule) error {
	rms_json_bytes, _ := json.Marshal(modules)

	// Compress
	rms_compressed, err := CompressBytes(rms_json_bytes)
	if err != nil {
		return err
	}

	ctx := context.Background()
	return cr.redis.Set(ctx, generateAllModulesForStationCacheKey(station_id), rms_compressed, 0).Err()
}

func (cr *CacheRepository) ClearAllModulesForStation(station_id string) error {
	ctx := context.Background()
	return cr.redis.Del(ctx, generateAllModulesForStationCacheKey(station_id)).Err()
}
