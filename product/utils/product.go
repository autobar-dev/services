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

	badges := []types.ProductBadge{}
	err = json.Unmarshal([]byte(pp.Badges), &badges)
	if err != nil {
		fmt.Printf("IMPORTANT: failed to parse badges for product: %v\n", pp)
	}

	return &types.Product{
		Id:           pp.Id,
		Names:        names_map,
		Descriptions: descriptions_map,
		Cover:        cover,
		Enabled:      pp.Enabled,
		Badges:       badges,
		CreatedAt:    pp.CreatedAt,
		UpdatedAt:    pp.UpdatedAt,
	}
}

func RedisProductToProduct(rp repositories.RedisProduct, cover filerepository.File) *types.Product {
	pbs := []types.ProductBadge{}
	for _, pb := range rp.Badges {
		pbs = append(pbs, *RedisProductBadgeToProductBadge(pb))
	}

	return &types.Product{
		Id:           rp.Id,
		Names:        rp.Names,
		Descriptions: rp.Descriptions,
		Cover:        cover,
		Enabled:      rp.Enabled,
		Badges:       pbs,
		CreatedAt:    rp.CreatedAt,
		UpdatedAt:    rp.UpdatedAt,
	}
}

func ProductBadgeToRedisProductBadge(pb types.ProductBadge) *repositories.RedisProductBadge {
	return &repositories.RedisProductBadge{
		Type:  repositories.RedisProductBadgeType(pb.Type),
		Label: pb.Label,
		Value: pb.Value,
	}
}

func RedisProductBadgeToProductBadge(rpb repositories.RedisProductBadge) *types.ProductBadge {
	return &types.ProductBadge{
		Type:  types.ProductBadgeType(rpb.Type),
		Label: rpb.Label,
		Value: rpb.Value,
	}
}

func ProductBadgeToMeiliProductBadge(pb types.ProductBadge) *repositories.MeiliProductBadge {
	return &repositories.MeiliProductBadge{
		Type:  repositories.MeiliProductBadgeType(pb.Type),
		Label: pb.Label,
		Value: pb.Value,
	}
}

func MeiliProductBadgeToProductBadge(mpb repositories.MeiliProductBadge) *types.ProductBadge {
	return &types.ProductBadge{
		Type:  types.ProductBadgeType(mpb.Type),
		Label: mpb.Label,
		Value: mpb.Value,
	}
}

func ProductToRedisProduct(p types.Product) *repositories.RedisProduct {
	rpbs := []repositories.RedisProductBadge{}
	for _, pb := range p.Badges {
		rpbs = append(rpbs, *ProductBadgeToRedisProductBadge(pb))
	}

	return &repositories.RedisProduct{
		Id:           p.Id,
		Names:        p.Names,
		Descriptions: p.Descriptions,
		CoverId:      p.Cover.Id,
		Enabled:      p.Enabled,
		Badges:       rpbs,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}
}

func MeiliProductToProduct(mp repositories.MeiliProduct, cover filerepository.File) *types.Product {
	mpbs := []types.ProductBadge{}
	for _, pb := range mp.Badges {
		mpbs = append(mpbs, *MeiliProductBadgeToProductBadge(pb))
	}

	return &types.Product{
		Id:           mp.Id,
		Names:        mp.Names,
		Descriptions: mp.Descriptions,
		Cover:        cover,
		Enabled:      mp.Enabled,
		Badges:       mpbs,
		CreatedAt:    mp.CreatedAt,
		UpdatedAt:    mp.UpdatedAt,
	}
}
