package types

import (
	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
)

type Meta struct {
	Hash    string
	Version string
}

type AppContext struct {
	Log  *log.Logger
	Meta *Meta
}

type RestContext struct {
	echo.Context
	*AppContext
}
