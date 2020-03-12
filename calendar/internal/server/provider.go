package server

import (
	"github.com/karasunokami/go.otus.hw/calendar/internal/app"
	"github.com/karasunokami/go.otus.hw/calendar/internal/config"
)

type Provider struct{}

func (p *Provider) Register(c *app.Container) {
	c.Set("server", func(c *app.Container) interface{} {
		conf := c.Get("config").(*config.AppConfig)

		return NewHttpServer(conf.HttpListen)
	})
}
