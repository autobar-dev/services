package controllers

import (
	"errors"
	"fmt"

	"github.com/autobar-dev/services/product/repositories"
	"github.com/autobar-dev/services/product/types"
	"github.com/autobar-dev/services/product/utils"
)

func CreateProduct(
	ac *types.AppContext,
	slug string,
	names map[string]string,
	descriptions map[string]string,
	cover_id string,
	enabled bool,
) error {
	pr := ac.Repositories.Product
	shr := ac.Repositories.SlugHistory
	mr := ac.Repositories.Meili
	cr := ac.Repositories.Cache

	product_result, _ := GetProductBySlug(ac, slug)
	if product_result != nil {
		return errors.New("product already exists")
	}

	product_id, err := pr.Create(names, descriptions, cover_id, enabled)
	if err != nil {
		return err
	}

	err = shr.Create(*product_id, slug)
	if err != nil {
		return err
	}

	product, err := GetProductById(ac, *product_id)
	if err != nil {
		fmt.Printf("failed to fetch newly created product: %v\n", err)
	}

	mpbs := []repositories.MeiliProductBadge{}
	for _, product_badge := range product.Badges {
		mpbs = append(mpbs, *utils.ProductBadgeToMeiliProductBadge(product_badge))
	}

	// Clear all products cache
	err = cr.ClearAllProducts()
	if err != nil {
		fmt.Printf("failed to clear cache for all products: %v", err)
	}

	err = mr.AddProduct(
		product.Id,
		product.Names,
		product.Descriptions,
		product.Cover.Id,
		product.Enabled,
		mpbs,
		product.CreatedAt,
		product.UpdatedAt,
	)

	return err
}
