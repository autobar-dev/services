package routes

import (
	"net/http"

	"github.com/autobar-dev/services/modulerealtime/types"
	"github.com/labstack/echo/v4"
)

type MetaResponse struct {
	Hash    string `json:"hash"`
	Version string `json:"version"`
}

func Meta(c echo.Context) error {
	app_ctx := c.(*types.RestContext)

	return app_ctx.JSON(http.StatusOK, &MetaResponse{
		Hash:    app_ctx.Meta.Hash,
		Version: app_ctx.Meta.Version,
	})
}
