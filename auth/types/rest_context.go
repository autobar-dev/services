package types

import (
	"github.com/labstack/echo/v4"
)

type RestContext struct {
	echo.Context
	AppContext    *AppContext
	ClientContext *ClientContext
}
