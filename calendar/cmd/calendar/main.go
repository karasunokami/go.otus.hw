package main

import (
	"fmt"
	"github.com/karasunokami/go.otus.hw/calendar/internal/app"
	"github.com/karasunokami/go.otus.hw/calendar/internal/config"
	"github.com/karasunokami/go.otus.hw/calendar/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	cont := app.NewContainer()

	cont.Set("logger", func(c *app.Container) interface{} {
		return new(logrus.FieldLogger)
	})

	cont.Register(new(config.Provider))
	cont.Register(new(server.Provider))

	err := cont.Get("server").(server.HttpServer).Run()
	fmt.Print(err)
}
