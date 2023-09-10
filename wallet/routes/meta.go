package routes

import (
	"github.com/autobar-dev/services/wallet/types"
	"github.com/autobar-dev/services/wallet/utils"
	"github.com/labstack/echo/v4"
)

type MetaRouteResponse struct {
	Status string      `json:"status"`
	Data   *types.Meta `json:"data"`
}

func MetaRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	return rest_context.JSON(200, &MetaRouteResponse{
		Status: "ok",
		Data:   utils.MetaFactorsToMeta(*app_context.MetaFactors),
	})
}
