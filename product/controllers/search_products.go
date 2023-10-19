package controllers

import (
	"github.com/autobar-dev/services/product/repositories"
	"github.com/autobar-dev/services/product/types"
	"github.com/autobar-dev/services/product/utils"
)

func SearchProducts(
	ac *types.AppContext,
	query string,
	hits_per_page int,
	page int,
	include_disabled bool,
) (*[]types.Product, error) {
	mr := ac.Repositories.Meili

	mps, err := mr.SearchProducts(&repositories.MeiliProductsSearchOptions{
		Query:           query,
		HitsPerPage:     hits_per_page,
		Page:            page,
		IncludeDisabled: include_disabled,
	})
	if err != nil {
		return nil, err
	}

	products := []types.Product{}
	for _, mp := range *mps {
		cover_file, err := ac.Repositories.File.GetFile(mp.CoverId)
		if err != nil {
			return nil, err
		}

		p := utils.MeiliProductToProduct(mp, *cover_file)
		products = append(products, *p)
	}

	return &products, nil
}
