package oauth

import (
	"github.com/autobar-dev/services/auth/types"
	"github.com/labstack/echo/v4"
)

type TokenResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
}

func TokenRoute(c echo.Context) error {
	rc := c.(*types.RestContext)
	ac := rc.AppContext

	oa_s := ac.Repositories.OAuthServer

	w := rc.Response().Writer
	r := rc.Request()

	err := oa_s.HandleTokenRequest(w, r)
	if err != nil {
		err := "incorrect token request"
		return rc.JSON(400, &AuthorizeResponse{
			Status: "error",
			Error:  &err,
		})
	}

	resp_err := "unknown internal error"
	return rc.JSON(500, &AuthorizeResponse{
		Status: "error",
		Error:  &resp_err,
	})
}
