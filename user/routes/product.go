package routes

import (
	"github.com/autobar-dev/services/product/controllers"
	"github.com/autobar-dev/services/product/types"
	"github.com/labstack/echo/v4"
)

type GetProductRouteResponseData struct {
	Type         types.ProductResponseType `json:"type"`
	Product      *types.Product            `json:"product"`
	RedirectSlug *string                   `json:"redirect_slug"`
}

type GetProductRouteResponse struct {
	Status string                       `json:"status"`
	Error  *string                      `json:"error"`
	Data   *GetProductRouteResponseData `json:"data"`
}

func GetProductRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	id := c.QueryParam("id")
	slug := c.QueryParam("slug")

	// Neither id nor slug specified
	if id == "" && slug == "" {
		err := "either id or slug must be provided"
		return rest_context.JSON(400, &GetProductRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	if id != "" { // Id specified
		product, err := controllers.GetProductById(&app_context, id)
		if err != nil {
			err := err.Error()
			return c.JSON(400, &GetProductRouteResponse{
				Status: "error",
				Error:  &err,
				Data:   nil,
			})
		}

		return c.JSON(200, &GetProductRouteResponse{
			Status: "ok",
			Error:  nil,
			Data: &GetProductRouteResponseData{
				Type:         types.DataProductResponseType,
				Product:      product,
				RedirectSlug: nil,
			},
		})
	}

	// Slug specified
	product_slug_response, err := controllers.GetProductBySlug(&app_context, slug)
	if err != nil {
		err := err.Error()
		return c.JSON(400, &GetProductRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	return rest_context.JSON(200, &GetProductRouteResponse{
		Status: "ok",
		Error:  nil,
		Data: &GetProductRouteResponseData{
			Type:         product_slug_response.Type,
			Product:      product_slug_response.Product,
			RedirectSlug: product_slug_response.RedirectSlug,
		},
	})
}
