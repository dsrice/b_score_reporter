package main

import (
	"go.uber.org/dig"
	"login/controllers"
	"login/infra/server"
)

func main() {
	c := dig.New()
	c.Provide(controllers.NewtestController)
	c.Provide(server.NewServer)
	c.Invoke(func(s server.Server) {
		s.Start()
	})
}
