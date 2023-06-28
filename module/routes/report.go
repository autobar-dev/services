package routes

import (
	"github.com/labstack/echo/v4"
	"go.a5r.dev/services/module/controllers"
	"go.a5r.dev/services/module/types"
)

type ReportRouteRequestBody struct {
	Queue string `json:"queue"`

	Status string `json:"status"`
}

type ReportRouteResponse struct {
	Status string  `json:"status"`
	Data   *string `json:"data"`
	Error  *string `json:"error"`
}

func ReportRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	var session_id *string

	session_from_query := c.QueryParam("session_id")
	session_from_cookie, _ := c.Cookie("session_id")

	if session_from_query != "" {
		session_id = &session_from_query
	} else if session_from_cookie != nil {
		session_id = &session_from_cookie.Value
	} else {
		err := "session not available from either session or cookie"

		return rest_context.JSON(400, &ReportRouteResponse{
			Status: "error",
			Data:   nil,
			Error:  &err,
		})
	}

	var rrrb ReportRouteRequestBody
	err := rest_context.Bind(&rrrb)
	if err != nil {
		err := "missing or incorrect values for queue or status body parameters"

		return rest_context.JSON(400, &ReportRouteResponse{
			Status: "error",
			Error:  &err,
			Data:   nil,
		})
	}

	session_data, err := controllers.VerifySessionController(&app_context, *session_id)
	if err != nil {
		err := err.Error()

		return rest_context.JSON(400, &ReportRouteResponse{
			Status: "error",
			Data:   nil,
			Error:  &err,
		})
	}

	if session_data.ClientType != types.ModuleClientType {
		err := "you cannot access this endpoint since client is not a module"
		return rest_context.JSON(400, &ReportRouteResponse{
			Status: "error",
			Data:   nil,
			Error:  &err,
		})
	}

	msr := &types.ModuleSentReport{
		Status: rrrb.Status,
	}

	err = controllers.ReportController(&app_context, rrrb.Queue, *msr)
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
