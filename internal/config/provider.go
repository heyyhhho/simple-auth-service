package config

import (
	"github.com/heyyhhho/simple-auth-service/internal/core"
	"github.com/heyyhhho/simple-auth-service/internal/logger"
	"github.com/sirupsen/logrus"
)

type Provider struct {
}

const (
	Key = "config"
)

func (p *Provider) Register(c core.Container) {
	p.registerConfig(c)
}

func (p *Provider) registerConfig(c core.Container) {
	loggerService := c.MustGet(logger.Key).(*logrus.Logger).WithField("service", Key)
	c.Set(Key, func(c core.Container) interface{} {
		return NewConfig(loggerService)
	})
}
