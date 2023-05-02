package adapters

import (
	"github.com/labstack/echo/v4"
	"login/domains/tables"
)

type UserAdapter interface {
	Create(user tables.User) error
	GetUser(sc tables.User) ([]*tables.User, error)
	SetContext(c echo.Context)
}
