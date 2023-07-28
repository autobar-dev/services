package controllers

import (
	"fmt"

	"github.com/autobar-dev/services/product/repositories"
	"github.com/autobar-dev/services/product/types"
	"github.com/autobar-dev/services/product/utils"
)

func EditProduct(ac *types.AppContext, id string, names map[string]string, descriptions map[string]string, cover string, enabled bool) (*[]string, error) {
	pr := ac.Repositories.Product
	// shr := ac.Repositories.SlugHistory
	// mr := ac.Repositories.Meili
	// cr := ac.Repositories.Cache
	pr.Edit(id, &repositories.PostgresEditProductInput{})

	// product, err := GetProductById(ac, id)
	// if err != nil {
	// 	return nil, err
	// }
	//
	// fields_altered := []string{}
	//
	// if !utils.CompareMaps(product.Descriptions, descriptions) {
	// 	fields_altered = append(fields_altered, "descriptions")
	// }
	// if !utils.CompareMaps(product.Names, names) {
	// 	fields_altered = append(fields_altered, "names")
	// }
	// if product.Cover != cover {
	// 	fields_altered = append(fields_altered, "cover")
	// }
	// if product.Enabled != enabled {
	// 	fields_altered = append(fields_altered, "enabled")
	// }
	//
	// fmt.Println("creating product slug in postgres")
	// err = shr.Create(*product_id, slug)
	// if err != nil {
	// 	return err
	// }
	//
	// product, err := GetProductById(ac, *product_id)
	// if err != nil {
	// 	fmt.Printf("failed to fetch newly created product: %v\n", err)
	// }
	//
	// // Clear all products cache
	// _ = cr.ClearAllProducts()
	//
	// fmt.Println("creating product in meili")
	// err = mr.AddProduct(product.Id, product.Names, product.Descriptions, product.Cover, product.Enabled, product.CreatedAt, product.UpdatedAt)
	//
	// return err
}
