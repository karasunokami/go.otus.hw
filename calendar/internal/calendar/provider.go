package calendar

import (
	"github.com/karasunokami/go.otus.hw/calendar/internal/app"
	"github.com/karasunokami/go.otus.hw/calendar/internal/db"
	"github.com/sirupsen/logrus"
)

type Provider struct{}

func (p *Provider) Register(c *app.Container) {
	c.Set("calendar", func(c *app.Container) interface{} {
		logger := c.Get("logger").(logrus.FieldLogger)
		client := c.Get("db.client").(db.Client)

		return NewCalendar(client, logger)
	})
}
