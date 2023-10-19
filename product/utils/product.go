package utils

import (
	"encoding/json"
	"fmt"

	"github.com/autobar-dev/services/product/repositories"
	"github.com/autobar-dev/services/product/types"
	filerepository "github.com/autobar-dev/shared-libraries/go/file-repository"
)

func PostgresProductToProduct(pp repositories.PostgresProduct, cover filerepository.File) *types.Product {
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
		Cover:        cover,
		Enabled:      pp.Enabled,
		CreatedAt:    pp.CreatedAt,
		UpdatedAt:    pp.UpdatedAt,
	}
}

func RedisProductToProduct(rp repositories.RedisProduct, cover filerepository.File) *types.Product {
	return &types.Product{
		Id:           rp.Id,
		Names:        rp.Names,
		Descriptions: rp.Descriptions,
		Cover:        cover,
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
		CoverId:      p.Cover.Id,
		Enabled:      p.Enabled,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}
}

func MeiliProductToProduct(mp repositories.MeiliProduct, cover filerepository.File) *types.Product {
	return &types.Product{
		Id:           mp.Id,
		Names:        mp.Names,
		Descriptions: mp.Descriptions,
		Cover:        cover,
		Enabled:      mp.Enabled,
		CreatedAt:    mp.CreatedAt,
		UpdatedAt:    mp.UpdatedAt,
	}
}
