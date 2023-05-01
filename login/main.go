package main

import (
	"github.com/labstack/gommon/log"
	"go.uber.org/dig"
	"login/controllers"
	"login/infra/server"
)

func main() {
	c := dig.New()
	err := c.Provide(controllers.NewtestController)

	if err != nil {
		log.Error("error")
	}

	err = c.Provide(server.NewServer)

	if err != nil {
		log.Error("error")
	}

	err = c.Invoke(func(s server.Server) {
		s.Start()
	})

	if err != nil {
		log.Error("error")
	}
}
