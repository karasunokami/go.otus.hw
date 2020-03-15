package logger

import (
	"github.com/karasunokami/go.otus.hw/calendar/internal/app"
	"github.com/karasunokami/go.otus.hw/calendar/internal/config"
	"github.com/sirupsen/logrus"
	"os"
)

type Provider struct{}

func (p *Provider) Register(c *app.Container) {
	c.Set("logger", func(c *app.Container) interface{} {
		conf := c.Get("config").(*config.AppConfig)
		logger := logrus.New()

		lvl, err := logrus.ParseLevel(conf.LogLevel)
		if err != nil {
			logger.Errorf("Failed to parse log level from config: %s", conf.LogLevel)
			logger.SetLevel(logrus.DebugLevel)
		} else {
			logger.SetLevel(lvl)
		}

		file, err := os.OpenFile(conf.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			logger.Out = file
		} else {
			logger.Errorf("Failed to log to file \"%s\", using default stderr", conf.LogFile)
		}

		return logger
	})
}
