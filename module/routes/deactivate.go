package routes

import (
	"github.com/autobar-dev/services/module/controllers"
	"github.com/autobar-dev/services/module/types"
	authrepository "github.com/autobar-dev/shared-libraries/go/auth-repository"
	"github.com/labstack/echo/v4"
)

type DeactivateRouteResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
}

func DeactivateRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := rest_context.AppContext
	client_context := rest_context.ClientContext

	if client_context == nil {
		err := "not authorized"
		return rest_context.JSON(401, &DeactivateRouteResponse{
			Status: "error",
			Error:  &err,
		})
	}

	if client_context.Type == authrepository.UserTokenOwnerType {
		err := controllers.DeactivateAsUserController(app_context, client_context.Identifier)
		if err != nil {
			err := err.Error()
			return rest_context.JSON(400, &DeactivateRouteResponse{
				Status: "error",
				Error:  &err,
			})
		}
	} else if client_context.Type == authrepository.ModuleTokenOwnerType {
		err := controllers.DeactivateAsModuleController(app_context, client_context.Identifier)
		if err != nil {
			err := err.Error()
			return rest_context.JSON(400, &DeactivateRouteResponse{
				Status: "error",
				Error:  &err,
			})
		}
	}

	return rest_context.JSON(200, &DeactivateRouteResponse{
		Status: "ok",
		Error:  nil,
	})
}
