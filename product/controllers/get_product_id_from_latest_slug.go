package controllers

import (
	"fmt"

	"github.com/autobar-dev/services/product/types"
)

func GetProductIdFromLatestSlug(ac *types.AppContext, latest_slug string) (*string, error) {
	cr := ac.Repositories.Cache
	shr := ac.Repositories.SlugHistory

	cached_product_id, _ := cr.GetProductIdFromLatestSlug(latest_slug)
	if cached_product_id != nil {
		return cached_product_id, nil
	}

	slug_entry, err := shr.Get(latest_slug)
	if err != nil {
		return nil, err
	}

	err = cr.SetProductIdFromLatestSlug(slug_entry.Slug, slug_entry.ProductId)
	if err != nil {
		fmt.Printf("failed to set cache for latest_slug->product_id: %v\n", err)
	}

	product_id := slug_entry.ProductId

	return &product_id, nil
}
