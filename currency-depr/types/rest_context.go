package types

import (
	"github.com/labstack/echo"
)

type RestContext struct {
	echo.Context
	*AppContext
}
