package server

import (
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"login/controllers/handlers"
)

type Server struct {
	TestController handlers.TestHandler
}

func NewServer(test handlers.TestHandler) Server {
	return Server{
		TestController: test,
	}
}

// NewServer
// サーバー起動処理
func (s *Server) Start() {
	e := echo.New()

	e.GET("/", s.TestController.Get)

	c := jaegertracing.New(e, nil)
	defer c.Close()

	e.Logger.Fatal(e.Start(":1323"))
}
