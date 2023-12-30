package middlewares

import (
	"errors"
	"net/http"
	"strings"
	"uploader-service/crash"
	authService "uploader-service/services/auth"

	"github.com/labstack/echo/v4"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorizationHeader := c.Request().Header.Get("Authorization")
		if authorizationHeader == "" {
			err := errors.New("Authorization header format not valid")
			return c.JSON(http.StatusUnauthorized, crash.GenerateError(crash.AuthHeaderNotValid, err))
		}

		authHeaderParts := strings.Split(authorizationHeader, " ")
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			err := errors.New("Authorization header format not valid")
			return c.JSON(http.StatusUnauthorized, crash.GenerateError(crash.AuthHeaderNotValid, err))
		}

		_, verifiedToken := authService.VerifyToken(authHeaderParts[1])
		if !verifiedToken {
			err := errors.New("Not valid token")
			return c.JSON(http.StatusForbidden, crash.GenerateError(crash.AuthHeaderNotValid, err))
		}

		return next(c)
	}
}
