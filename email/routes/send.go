package routes

import (
	"github.com/autobar-dev/services/email/controllers"
	"github.com/autobar-dev/services/email/types"
	"github.com/labstack/echo/v4"
)

type SendRequestBodyMessage struct {
	Plain string `json:"plain"`
	Html  string `json:"html"`
}

type SendRequestBody struct {
	From    string                 `json:"from"`
	To      string                 `json:"to"`
	Subject string                 `json:"subject"`
	Message SendRequestBodyMessage `json:"message"`
}

type SendResponseBody struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
}

func SendRoute(c echo.Context) error {
	rest_context := c.(*types.RestContext)
	app_context := *(*rest_context).AppContext

	var srb SendRequestBody
	err := rest_context.Bind(&srb)
	if err != nil {
		err := "invalid request body"
		return rest_context.JSON(400, &SendResponseBody{
			Status: "error",
			Error:  &err,
		})
	}

	err = controllers.Send(&app_context, srb.From, srb.To, srb.Subject, srb.Message.Plain, srb.Message.Html)
	if err != nil {
		return rest_context.JSON(400, &SendResponseBody{
			Status: "error",
			Error:  nil,
		})
	}

	return rest_context.JSON(200, &SendResponseBody{
		Status: "ok",
		Error:  nil,
	})
}
