package handlers

import "github.com/labstack/echo/v4"

type TestHandler interface {
	Get(c echo.Context) error
}
