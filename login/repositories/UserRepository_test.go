package repositories_test

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"login/domains/tables"
	"login/repositories"
	"login/repositories/adapters"
	"login/test"
	"net/http"
	"net/http/httptest"
	"testing"
)

type GetUserSuite struct {
	suite.Suite
	repo adapters.UserAdapter
}

func (s *GetUserSuite) SetupTest() {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c := e.NewContext(req, rec)

	db := test.DataBase()

	s.repo = repositories.NewUser(db)
	s.repo.SetContext(c)
}

func (s *GetUserSuite) TestSuccess() {
	sc := tables.User{}

	userList, err := s.repo.GetUser(sc)

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), userList)
}

func TestRepositoryGetUserSuite(t *testing.T) {
	suite.Run(t, new(GetUserSuite))
}
