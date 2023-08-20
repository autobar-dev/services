package routes

import (
	"github.com/autobar-dev/services/user/controllers"
	"github.com/autobar-dev/services/user/types"
	"github.com/labstack/echo/v4"
)

type CreateUserRequestBody struct {
	Email string `json:"email"`

}

type CreateProductRouteResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
}

func CreateProductRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	var cprb CreateUserRequestBody

	err := c.Bind(&cprb)
	if err != nil {
		err := err.Error()
		return c.JSON(400, &CreateProductRouteResponse{
			Status: "error",
			Error:  &err,
		})
	}

	if len(cprb.Names) == 0 || cprb.Slug == "" {
		err := "either slug or names empty"
		return c.JSON(400, &CreateProductRouteResponse{
			Status: "error",
			Error:  &err,
		})
	}

	err = controllers.CreateProduct(&app_context, cprb.Slug, cprb.Names, cprb.Descriptions, cprb.Cover, true)
	if err != nil {
		err := err.Error()
		return c.JSON(400, &CreateProductRouteResponse{
			Status: "error",
			Error:  &err,
		})
	}

	return c.JSON(200, &CreateProductRouteResponse{
		Status: "ok",
		Error:  nil,
	})
}
