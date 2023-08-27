package middleware

import (
	"regexp"

	"github.com/autobar-dev/services/auth/types"
	"github.com/autobar-dev/shared-libraries/go/auth-repository"
	"github.com/labstack/echo/v4"
)

func AccessTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		rest_context := c.(*types.RestContext)
		app_context := rest_context.AppContext

		authorization_header := c.Request().Header.Get("Authorization")

		pattern := `Bearer\s+(\S+)`
		re := regexp.MustCompile(pattern)

		rest_context.ClientContext = nil

		submatches := re.FindStringSubmatch(authorization_header)
		if len(submatches) == 2 {
			at_string := submatches[1]
			at_payload, err := authrepository.ParseAccessToken(app_context.Config.JwtSecret, at_string)
			if err == nil {
				rest_context.ClientContext = &types.ClientContext{
					Type:       at_payload.ClientType,
					Identifier: at_payload.Identifier,
					Role:       at_payload.Role,
				}
			}
		}
		return next(c)
	}
}
