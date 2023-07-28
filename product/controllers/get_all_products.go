package controllers

import (
	"fmt"

	"github.com/autobar-dev/services/product/repositories"
	"github.com/autobar-dev/services/product/types"
	"github.com/autobar-dev/services/product/utils"
)

func GetAllProducts(ac *types.AppContext) (*[]types.Product, error) {
	pr := ac.Repositories.Product
	cr := ac.Repositories.Cache

	// Try cache first
	rps, err := cr.GetAllProducts()
	if rps != nil {
		products := []types.Product{}
		for _, rp := range *rps {
			products = append(products, *utils.RedisProductToProduct(rp))
		}

		return &products, nil
	}

	// Use database otherwise
	pps, err := pr.GetAll()
	if err != nil {
		return nil, err
	}

	products := []types.Product{}

	for _, pp := range *pps {
		product := utils.PostgresProductToProduct(pp)

		// Set cache
		_ = cr.SetProduct(product.Id, product.Names, product.Descriptions, product.Cover, product.Enabled, product.CreatedAt, product.UpdatedAt)

		products = append(products, *utils.PostgresProductToProduct(pp))
	}

	// Set cache for all products
	rps_save := []repositories.RedisProduct{}
	for _, p := range products {
		rps_save = append(rps_save, *utils.ProductToRedisProduct(p))
	}
	_ = cr.SetAllProducts(rps_save)

	return &products, nil
}
