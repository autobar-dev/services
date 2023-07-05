package types

import (
	echo "github.com/labstack/echo/v4"
)

type RestContext struct {
	echo.Context
	AppContext *AppContext
}
