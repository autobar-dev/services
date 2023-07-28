package controllers

import (
	"errors"
	"fmt"

	"github.com/autobar-dev/services/product/repositories"
	"github.com/autobar-dev/services/product/types"
	"github.com/autobar-dev/services/product/utils"
)

func EditProduct(ac *types.AppContext, id string, slug *string, names *map[string]string, descriptions *map[string]string, cover *string, enabled *bool) (*[]string, error) {
	pr := ac.Repositories.Product
	shr := ac.Repositories.SlugHistory
	mr := ac.Repositories.Meili
	cr := ac.Repositories.Cache

	product, err := GetProductById(ac, id)
	if err != nil {
		return nil, err
	}

	fields_altered := []string{}

	// Check if slug already exists
	if slug != nil {
		_, err = GetLatestSlugFromSlug(ac, *slug)
		if err == nil { // if getting latest slug doesn't error out, it means that it already exists
			return nil, errors.New("the provided slug already exists")
		}

		// get all slugs for product
		slug_entries, err := shr.GetAllSlugsForProduct(id)
		if err != nil {
			return nil, err
		}

		slugs := []string{}
		for _, pshe := range *slug_entries {
			slugs = append(slugs, pshe.Slug)
		}

		// get latest slug for product
		latest_slug, err := GetLatestSlugFromSlug(ac, slugs[0])
		if err != nil {
			return nil, err
		}

		fields_altered = append(fields_altered, "slug")

		err = shr.Create(id, *slug)
		if err != nil {
			return nil, err
		}

		err = cr.ClearProductIdFromLatestSlug(*latest_slug)
		if err != nil {
			fmt.Printf("failed to clear product from latest slug cache: %v\n", err)
		}

		err = cr.ClearMultipleSlugsToLatestSlug(slugs)
		if err != nil {
			fmt.Printf("failed to clear multiple slugs to latest slug cache: %v\n", err)
		}
	}

	if names != nil && !utils.CompareMaps(product.Names, *names) {
		fields_altered = append(fields_altered, "names")
	}
	if descriptions != nil && !utils.CompareMaps(product.Descriptions, *descriptions) {
		fields_altered = append(fields_altered, "descriptions")
	}
	if cover != nil && product.Cover != *cover {
		fields_altered = append(fields_altered, "cover")
	}
	if enabled != nil && product.Enabled != *enabled {
		fields_altered = append(fields_altered, "enabled")
	}

	err = pr.Edit(id, &repositories.PostgresEditProductInput{
		Names:        names,
		Descriptions: descriptions,
		Cover:        cover,
		Enabled:      enabled,
	})
	if err != nil {
		return nil, err
	}

	// Clear cache
	err = cr.ClearProduct(id)
	if err != nil {
		fmt.Printf("failed to clear cache for product: %v", err)
	}

	err = cr.ClearAllProducts()
	if err != nil {
		fmt.Printf("failed to clear cache for all products: %v", err)
	}

	// Fetch the product
	product, err = GetProductById(ac, id)
	if err != nil {
		return nil, err
	}

	err = mr.AddProduct(product.Id, product.Names, product.Descriptions, product.Cover, product.Enabled, product.CreatedAt, product.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &fields_altered, nil
}
