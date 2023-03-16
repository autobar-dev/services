package types

import (
	"github.com/autobar-dev/services/currency/types/interfaces"
	"github.com/labstack/echo"
)

type RestContext struct {
	echo.Context

	AppLogger *interfaces.AppLogger
	Stores    *AppStores
}
