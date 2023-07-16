package controllers

import (
	"github.com/autobar-dev/services/product/types"
)

type GetProductBySlugResult struct {
	Type         types.ProductResponseType
	Product      *types.Product
	RedirectSlug *string
}

func GetProductBySlug(ac *types.AppContext, slug string) (*GetProductBySlugResult, error) {
	latest_slug, err := GetLatestSlugFromSlug(ac, slug)
	if err != nil {
		return nil, err
	}

	if slug != *latest_slug {
		return &GetProductBySlugResult{
			Type:         types.RedirectProductResponseType,
			Product:      nil,
			RedirectSlug: latest_slug,
		}, nil
	}

	product_id, err := GetProductIdFromLatestSlug(ac, *latest_slug)
	if err != nil {
		return nil, err
	}

	product, err := GetProductById(ac, *product_id)
	if err != nil {
		return nil, err
	}

	return &GetProductBySlugResult{
		Type:         types.DataProductResponseType,
		Product:      product,
		RedirectSlug: nil,
	}, nil
}
