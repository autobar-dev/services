package routes

import (
	"github.com/autobar-dev/services/module/controllers"
	"github.com/autobar-dev/services/module/types"
	authrepository "github.com/autobar-dev/shared-libraries/go/auth-repository"
	"github.com/labstack/echo/v4"
)

type PrepareModuleRouteResponse struct {
	Status string                   `json:"status"`
	Data   *types.PrepareModuleData `json:"data"`
	Error  *string                  `json:"error"`
}

func PrepareModuleRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := rest_context.AppContext
	client_context := rest_context.ClientContext

	if client_context == nil {
		err := "not authorized"
		return rest_context.JSON(400, &PrepareModuleRouteResponse{
			Status: "error",
			Data:   nil,
			Error:  &err,
		})
	}

	if client_context.Type != authrepository.ModuleTokenOwnerType {
		err := "you are not a module"
		return rest_context.JSON(401, &PrepareModuleRouteResponse{
			Status: "error",
			Data:   nil,
			Error:  &err,
		})
	}

	pmd, err := controllers.PrepareModuleController(app_context, client_context.Identifier)
	if err != nil {
		err := err.Error()

		return rest_context.JSON(400, &PrepareModuleRouteResponse{
			Status: "error",
			Data:   nil,
			Error:  &err,
		})
	}

	return rest_context.JSON(200, &PrepareModuleRouteResponse{
		Status: "ok",
		Data:   pmd,
		Error:  nil,
	})
}
