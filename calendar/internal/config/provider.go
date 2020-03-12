package config

import (
	goflag "flag"
	"github.com/karasunokami/go.otus.hw/calendar/internal/app"
	flag "github.com/spf13/pflag"
)

type Provider struct{}

func (p *Provider) Register(c *app.Container) {
	c.Set("config.builder", func(c *app.Container) interface{} {
		return NewBuilder()
	})

	c.Set("config", func(c *app.Container) interface{} {
		builder := c.Get("config.builder").(Builder)

		var path = flag.String("config", "configs/config.yaml", "path to config file")
		flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
		flag.Parse()

		conf := builder.Build(*path)

		return conf
	})
}
