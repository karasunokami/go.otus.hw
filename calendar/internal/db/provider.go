package db

import (
	"github.com/karasunokami/go.otus.hw/calendar/internal/app"
)

type Provider struct{}

func (p *Provider) Register(c *app.Container) {
	c.Set("db.client", func(c *app.Container) interface{} {
		return NewClient()
	})
}
