package routes

import (
	"github.com/autobar-dev/services/product/controllers"
	"github.com/autobar-dev/services/product/types"
	"github.com/labstack/echo/v4"
)

type SearchProductsRouteResponseData struct {
	Hits *[]types.Product `json:"hits"`
}

type SearchProductsRouteResponse struct {
	Status string                           `json:"status"`
	Error  *string                          `json:"error"`
	Data   *SearchProductsRouteResponseData `json:"data"`
}

type SearchProductsRouteRequestBody struct {
	Query           string `json:"query"`
	HitsPerPage     int    `json:"hits_per_page"`
	Page            int    `json:"page"`
	IncludeDisabled bool   `json:"include_disabled"`
}

func SearchProductsRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	var sprb SearchProductsRouteRequestBody
	err := rest_context.Bind(&sprb)
	if err != nil {
		err := "query, hits_per_page and page request body parameters have to be provided"
		return rest_context.JSON(400, &SearchProductsRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	products, err := controllers.SearchProducts(&app_context, sprb.Query, sprb.HitsPerPage, sprb.Page, sprb.IncludeDisabled)
	if err != nil {
		err := "failed to perform search"
		return rest_context.JSON(500, &SearchProductsRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	return rest_context.JSON(200, &SearchProductsRouteResponse{
		Status: "ok",
		Error:  nil,
		Data: &SearchProductsRouteResponseData{
			Hits: products,
		},
	})
}
