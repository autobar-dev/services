package controllers

import (
	"fmt"

	"github.com/autobar-dev/services/user/types"
	"github.com/autobar-dev/services/user/utils"
)

func GetUserById(ac *types.AppContext, id string) (*types.User, error) {
	cr := *ac.Repositories.Cache
	ur := *ac.Repositories.User

	ru, _ := cr.GetUser(id)
	if ru != nil {
		return utils.RedisUserToUser(*ru), nil
	}

	pp, err := ur.Get(id)
	if err != nil {
		return nil, err
	}

	product := utils.PostgresUserToUser(*pp)

	err = cr.SetProduct(product.Id, product.Names, product.Descriptions, product.Cover, product.Enabled, product.CreatedAt, product.UpdatedAt)
	if err != nil {
		fmt.Printf("failed to set cache for product_id->product: %v\n", err)
	}

	return product, nil
}
