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
	fr := ac.Repositories.File

	// Try cache first
	rps, err := cr.GetAllProducts()
	if rps != nil {
		products := []types.Product{}
		for _, rp := range *rps {
			cover_file, err := fr.GetFile(rp.CoverId)
			if err != nil {
				fmt.Printf("(from cache) failed to get cover file when getting all products: %v", err)
				return nil, err
			}

			products = append(products, *utils.RedisProductToProduct(rp, *cover_file))
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
		cover_file, err := fr.GetFile(pp.CoverId)
		if err != nil {
			fmt.Printf("failed to get cover file when getting all products: %v", err)
			return nil, err
		}

		product := utils.PostgresProductToProduct(pp, *cover_file)

		// Set cache
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
			fmt.Printf("failed to set cache when getting all products: %v", err)
		}

		products = append(products, *product)
	}

	// Set cache for all products
	rps_save := []repositories.RedisProduct{}
	for _, p := range products {
		rps_save = append(rps_save, *utils.ProductToRedisProduct(p))
	}
	fmt.Println("all products cache miss. updating...")
	err = cr.SetAllProducts(rps_save)
	if err != nil {
		fmt.Printf("failed to set cache when getting all products: %v", err)
	}

	return &products, nil
}
