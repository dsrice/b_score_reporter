package controllers

import (
	"github.com/labstack/echo/v4"
	"login/controllers/handlers"
	"net/http"
)

type testController struct {
}

func NewtestController() handlers.TestHandler {
	return &testController{}
}

func (ct *testController) Get(c echo.Context) error {
	return c.String(http.StatusOK, "test")
}
