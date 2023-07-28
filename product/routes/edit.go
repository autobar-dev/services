package routes

import (
	"github.com/autobar-dev/services/product/controllers"
	"github.com/autobar-dev/services/product/types"
	"github.com/labstack/echo/v4"
)

type EditProductRequestBody struct {
	Id           string             `json:"id"`
	Slug         *string            `json:"slug"`
	Names        *map[string]string `json:"names"`
	Descriptions *map[string]string `json:"descriptions"`
	Cover        *string            `json:"cover"`
	Enabled      *bool              `json:"enabled"`
}

type EditProductRouteResponse struct {
	Status string    `json:"status"`
	Error  *string   `json:"error"`
	Data   *[]string `json:"data"`
}

func EditProductRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	var eprb EditProductRequestBody

	err := c.Bind(&eprb)
	if err != nil {
		err := err.Error()
		return c.JSON(400, &EditProductRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	_, err = controllers.GetProductById(&app_context, eprb.Id)
	if err != nil {
		err := "failed to fetch product with the provided id"
		return c.JSON(400, &EditProductRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	fields_altered, err := controllers.EditProduct(&app_context, eprb.Id, eprb.Slug, eprb.Names, eprb.Descriptions, eprb.Cover, eprb.Enabled)
	if err != nil {
		err := err.Error()
		return c.JSON(400, &EditProductRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	return c.JSON(200, &EditProductRouteResponse{
		Status: "ok",
		Error:  nil,
		Data:   fields_altered,
	})
}
