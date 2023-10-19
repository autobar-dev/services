package controllers

import (
	"fmt"

	"github.com/autobar-dev/services/product/types"
	"github.com/autobar-dev/services/product/utils"
)

func GetProductById(ac *types.AppContext, id string) (*types.Product, error) {
	cr := *ac.Repositories.Cache
	pr := *ac.Repositories.Product
	fr := *ac.Repositories.File

	rp, _ := cr.GetProduct(id)
	if rp != nil {
		cover_file, err := fr.GetFile(rp.CoverId)
		if err != nil {
			fmt.Printf("(from cache) failed to get cover file when getting product by id: %v", err)
			return nil, err
		}

		return utils.RedisProductToProduct(*rp, *cover_file), nil
	}

	pp, err := pr.Get(id)
	if err != nil {
		return nil, err
	}

	cover_file, err := fr.GetFile(pp.CoverId)
	if err != nil {
		fmt.Printf("failed to get cover file when getting product by id: %v", err)
		return nil, err
	}

	product := utils.PostgresProductToProduct(*pp, *cover_file)

	err = cr.SetProduct(
		product.Id,
		product.Names,
		product.Descriptions,
		product.Cover.Id,
		product.Enabled,
		product.CreatedAt,
		product.UpdatedAt,
	)
	if err != nil {
		fmt.Printf("failed to set cache for product_id->product: %v\n", err)
	}

	return product, nil
}
