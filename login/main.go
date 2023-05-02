package main

import (
	"github.com/labstack/gommon/log"
	"go.uber.org/dig"
	"login/controllers"
	"login/infra/database"
	"login/infra/server"
	"login/repositories"
)

func main() {
	c := dig.New()
	err := c.Provide(database.NewDataBase)

	if err != nil {
		log.Error("error %w", err)
	}

	err = c.Provide(repositories.NewUser)

	if err != nil {
		log.Error("error %w", err)
	}

	err = c.Provide(controllers.NewtestController)

	if err != nil {
		log.Error("error %w", err)
	}

	err = c.Provide(server.NewServer)

	if err != nil {
		log.Error("error %w", err)
	}

	err = c.Invoke(func(s server.Server) {
		s.Start()
	})

	if err != nil {
		log.Error("error %w", err)
	}
}
