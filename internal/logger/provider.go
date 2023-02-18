package logger

import (
	"github.com/heyyhhho/simple-auth-service/internal/core"
	"github.com/sirupsen/logrus"
	"os"
)

type Provider struct {
}

const Key = "logger"

func (p *Provider) Register(c core.Container) {
	p.registerLogger(c)
}

func (p *Provider) registerLogger(c core.Container) {
	c.Set(Key, func(c core.Container) interface{} {
		logger := logrus.New()
		logger.Out = os.Stdout
		logger.Level = logrus.DebugLevel
		logger.SetFormatter(&logrus.JSONFormatter{})
		logger.Debugln("Logger service init success")
		return logger
	})
}
