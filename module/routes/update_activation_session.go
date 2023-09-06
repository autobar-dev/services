package routes

import (
	"github.com/autobar-dev/services/module/controllers"
	"github.com/autobar-dev/services/module/types"
	authrepository "github.com/autobar-dev/shared-libraries/go/auth-repository"
	"github.com/labstack/echo/v4"
)

type UpdateActivationSessionRouteRequestBody struct {
	Price             int `json:"price"`
	AmountMillilitres int `json:"amount_millilitres"`
}

type UpdateActivationSessionRouteResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
}

func UpdateActivationSessionRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := rest_context.AppContext
	client_context := rest_context.ClientContext

	if client_context == nil {
		err := "not authorized"
		return rest_context.JSON(401, &UpdateActivationSessionRouteResponse{
			Status: "error",
			Error:  &err,
		})
	}

	if client_context.Type != authrepository.ModuleTokenOwnerType {
		err := "you are not a module"
		return rest_context.JSON(401, &UpdateActivationSessionRouteResponse{
			Status: "error",
			Error:  &err,
		})
	}

	var uasrrb UpdateActivationSessionRouteRequestBody
	if err := c.Bind(&uasrrb); err != nil {
		err := err.Error()
		return rest_context.JSON(401, &UpdateActivationSessionRouteResponse{
			Status: "error",
			Error:  &err,
		})
	}

	serial_number := client_context.Identifier

	err := controllers.UpdateActivationSessionController(app_context, serial_number, uasrrb.Price, uasrrb.AmountMillilitres)
	if err != nil {
		err := err.Error()
		return rest_context.JSON(400, &UpdateActivationSessionRouteResponse{
			Status: "error",
			Error:  &err,
		})
	}

	return rest_context.JSON(200, &UpdateActivationSessionRouteResponse{
		Status: "ok",
		Error:  nil,
	})
}
