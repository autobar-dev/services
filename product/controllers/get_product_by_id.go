package controllers

import (
	"fmt"

	"github.com/autobar-dev/services/product/types"
	"github.com/autobar-dev/services/product/utils"
)

func GetProductById(ac *types.AppContext, id string) (*types.Product, error) {
	cr := *ac.Repositories.Cache
	pr := *ac.Repositories.Product

	rp, _ := cr.GetProduct(id)
	if rp != nil {
		return utils.RedisProductToProduct(*rp), nil
	}

	pp, err := pr.Get(id)
	if err != nil {
		return nil, err
	}

	product := utils.PostgresProductToProduct(*pp)

	err = cr.SetProduct(product.Id, product.Names, product.Descriptions, product.Cover, product.Enabled, product.CreatedAt, product.UpdatedAt)
	if err != nil {
		fmt.Printf("failed to set cache for product_id->product: %v\n", err)
	}

	return product, nil
}
