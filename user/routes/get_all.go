package routes

import (
	"github.com/autobar-dev/services/product/controllers"
	"github.com/autobar-dev/services/product/types"
	"github.com/labstack/echo/v4"
)

type GetAllProductsRouteResponse struct {
	Status string           `json:"status"`
	Error  *string          `json:"error"`
	Data   *[]types.Product `json:"data"`
}

func GetAllProductsRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	products, err := controllers.GetAllProducts(&app_context)
	if err != nil {
		err := "failed to retrieve all products"
		return rest_context.JSON(500, &GetAllProductsRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	return rest_context.JSON(200, &GetAllProductsRouteResponse{
		Status: "ok",
		Error:  nil,
		Data:   products,
	})
}
