package controllers

import (
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go/log"
	"login/controllers/handlers"
	"login/domains/tables"
	"login/repositories/adapters"
	"net/http"
)

type testController struct {
	repo adapters.UserAdapter
}

func NewtestController(repo adapters.UserAdapter) handlers.TestHandler {
	return &testController{
		repo: repo,
	}
}

func (ct *testController) Get(c echo.Context) error {
	sp := jaegertracing.CreateChildSpan(c, "Controller")
	defer sp.Finish()

	sp.SetTag("testController", "Get")
	sp.LogFields(log.String("controller", "start"))

	ct.repo.SetContext(c)
	userList, err := ct.repo.GetUser(tables.User{})

	if err != nil {
		log.Error(err)
	}

	return c.JSON(http.StatusOK, userList)
}
