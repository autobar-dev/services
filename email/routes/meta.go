package routes

import (
	"github.com/autobar-dev/services/email/types"
	"github.com/autobar-dev/services/email/utils"

	"github.com/labstack/echo/v4"
)

type MetaRouteResponse struct {
	Status string      `json:"status"`
	Data   *types.Meta `json:"data"`
}

func MetaRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	meta := utils.GetMetaFromFactors(app_context.MetaFactors)

	return rest_context.JSON(200, &MetaRouteResponse{
		Status: "ok",
		Data:   meta,
	})
}
