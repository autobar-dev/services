package utils

import (
	"encoding/json"
	"fmt"

	"github.com/autobar-dev/services/product/repositories"
	"github.com/autobar-dev/services/product/types"
)

func PostgresProductToProduct(pp repositories.PostgresProduct) *types.Product {
	names_map := map[string]string{}
	err := json.Unmarshal([]byte(pp.Names), &names_map)
	if err != nil {
		fmt.Printf("IMPORTANT: failed to parse names for product: %v\n", pp)
	}

	descriptions_map := map[string]string{}
	err = json.Unmarshal([]byte(pp.Descriptions), &descriptions_map)
	if err != nil {
		fmt.Printf("IMPORTANT: failed to parse descriptions for product: %v\n", pp)
	}

	return &types.Product{
		Id:           pp.Id,
		Names:        names_map,
		Descriptions: descriptions_map,
		Cover:        pp.Cover,
		Enabled:      pp.Enabled,
		CreatedAt:    pp.CreatedAt,
		UpdatedAt:    pp.UpdatedAt,
	}
}

func RedisProductToProduct(rp repositories.RedisProduct) *types.Product {
	return &types.Product{
		Id:           rp.Id,
		Names:        rp.Names,
		Descriptions: rp.Descriptions,
		Cover:        rp.Cover,
		Enabled:      rp.Enabled,
		CreatedAt:    rp.CreatedAt,
		UpdatedAt:    rp.UpdatedAt,
	}
}

func ProductToRedisProduct(p types.Product) *repositories.RedisProduct {
	return &repositories.RedisProduct{
		Id:           p.Id,
		Names:        p.Names,
		Descriptions: p.Descriptions,
		Cover:        p.Cover,
		Enabled:      p.Enabled,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}
}

func MeiliProductToProduct(mp repositories.MeiliProduct) *types.Product {
	return &types.Product{
		Id:           mp.Id,
		Names:        mp.Names,
		Descriptions: mp.Descriptions,
		Cover:        mp.Cover,
		CreatedAt:    mp.CreatedAt,
		UpdatedAt:    mp.UpdatedAt,
	}
}
