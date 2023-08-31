package routes

import (
	"fmt"

	"github.com/autobar-dev/services/module/controllers"
	"github.com/autobar-dev/services/module/types"
	"github.com/autobar-dev/services/module/utils"
	"github.com/labstack/echo/v4"
)

type ReportRouteRequestBody struct {
	Queue  string `json:"queue"`
	Status string `json:"status"`
}

type ReportRouteResponse struct {
	Status string  `json:"status"`
	Data   *string `json:"data"`
	Error  *string `json:"error"`
}

func ReportRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := rest_context.AppContext
	client_context := rest_context.ClientContext

	fmt.Printf("request on /report\n")

	var rrrb ReportRouteRequestBody
	err := rest_context.Bind(&rrrb)
	if err != nil {
		fmt.Println(err)
		err := "missing or incorrect values for queue or status body parameters"

		return rest_context.JSON(400, &ReportRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	if client_context == nil {
		err := "not logged in"
		return rest_context.JSON(400, &ReportRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	client_type := *utils.ServiceClientTypeToClientType(client_context.Type)

	if client_type != types.ModuleClientType {
		err := "you cannot access this endpoint since client is not a module"
		return rest_context.JSON(400, &ReportRouteResponse{
			Status: "error",
			Data:   nil,
			Error:  &err,
		})
	}

	fmt.Printf("got report request from client type: %s, queue: %s\n", client_type, rrrb.Queue)

	msr := &types.ModuleSentReport{
		Status: rrrb.Status,
	}

	err = controllers.ReportController(app_context, rrrb.Queue, *msr)
	if err != nil {
		err := err.Error()

		return rest_context.JSON(400, &ReportRouteResponse{
			Status: "error",
			Data:   nil,
			Error:  &err,
		})
	}

	return rest_context.JSON(200, &CreateModuleRouteResponse{
		Status: "ok",
		Data:   nil,
		Error:  nil,
	})
}
