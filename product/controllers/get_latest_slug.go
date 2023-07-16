package controllers

import (
	"fmt"

	"github.com/autobar-dev/services/product/types"
)

func GetLatestSlugFromSlug(ac *types.AppContext, slug string) (*string, error) {
	cr := ac.Repositories.Cache
	shr := ac.Repositories.SlugHistory

	// Try cache first
	latest_slug, _ := cr.GetLatestSlugFromSlug(slug)
	if latest_slug != nil {
		return latest_slug, nil
	}

	fmt.Println("latest slug cache miss. updating...")

	// Cache miss. Calculate latest slug and update cache
	slug_entry, err := shr.Get(slug)
	if err != nil {
		return nil, err
	}

	product_id := slug_entry.ProductId
	all_slug_entries, err := shr.GetAllSlugsForProduct(product_id)
	if err != nil {
		return nil, err
	}

	ase := *all_slug_entries

	// Cache all slugs to latest slug
	all_slugs := []string{}
	for _, se := range ase {
		all_slugs = append(all_slugs, se.Slug)
	}

	err = cr.SetMultipleSlugsToLatestSlug(all_slugs)
	if err != nil {
		fmt.Printf("failed to set cache for slug->latest_slug: %v\n", err)
	}

	// Cache latest slug to product id
	latest_slug_entry := ase[len(ase)-1]
	err = cr.SetProductIdFromLatestSlug(latest_slug_entry.Slug, latest_slug_entry.ProductId)
	if err != nil {
		fmt.Printf("failed to set cache for latest_slug->product_id: %v\n", err)
	}

	return &latest_slug_entry.Slug, nil
}
