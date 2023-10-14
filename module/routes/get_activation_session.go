package routes

import (
	"github.com/autobar-dev/services/module/controllers"
	"github.com/autobar-dev/services/module/types"
	authrepository "github.com/autobar-dev/shared-libraries/go/auth-repository"
	"github.com/labstack/echo/v4"
)

type GetActivationSessionRouteResponse struct {
	Status string                   `json:"status"`
	Data   *types.ActivationSession `json:"data"`
	Error  *string                  `json:"error"`
}

func GetActivationSessionRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext
	client_context := rest_context.ClientContext

	if client_context == nil {
		err := "not authorized"
		return rest_context.JSON(401, &GetActivationSessionRouteResponse{
			Status: "error",
			Data:   nil,
			Error:  &err,
		})
	}

	if client_context.Type != authrepository.UserTokenOwnerType {
		err := "not a user"
		return rest_context.JSON(401, &GetActivationSessionRouteResponse{
			Status: "error",
			Data:   nil,
			Error:  &err,
		})
	}

	user_id := (*client_context).Identifier

	as, err := controllers.GetActivationSession(&app_context, user_id)
	if err != nil {
		err := err.Error()

		return rest_context.JSON(400, &GetActivationSessionRouteResponse{
			Status: "error",
			Data:   nil,
			Error:  &err,
		})
	}

	return rest_context.JSON(200, &GetActivationSessionRouteResponse{
		Status: "ok",
		Data:   as,
		Error:  nil,
	})
}
