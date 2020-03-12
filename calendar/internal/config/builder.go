package config

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Builder interface {
	Build(path string) *AppConfig
}

type builderImpl struct {
	logger logrus.FieldLogger

	config AppConfig
}

func NewBuilder() Builder {
	return &builderImpl{}
}

func (b *builderImpl) Build(path string) *AppConfig {
	b.config = AppConfig{
		HttpListen: HttpListen{
			Ip:   "127.0.0.1",
			Port: "1080",
		},
		LogFile:  "/var/log/calendar",
		LogLevel: "info",
	}

	logFields := logrus.Fields{
		"module": "config",
		"util":   "builder",
		"cmp":    "build",
	}

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		b.logger.WithFields(logFields).WithError(err).Error("failed to read config file")
	}

	err = yaml.Unmarshal(yamlFile, &b.config)

	if err != nil {
		b.logger.WithFields(logFields).WithError(err).Error("failed to parse config file")
	}

	return &b.config
}
