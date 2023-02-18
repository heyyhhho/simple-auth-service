package mysql

import (
	"github.com/heyyhhho/simple-auth-service/internal/config"
	"github.com/heyyhhho/simple-auth-service/internal/core"
	"github.com/heyyhhho/simple-auth-service/internal/logger"
	"github.com/sirupsen/logrus"
)

type Provider struct {
}

const (
	Key = "mysql"
)

func (p *Provider) Register(c core.Container) {
	p.registerMysql(c)
}

func (p *Provider) registerMysql(c core.Container) {
	appConfig := c.MustGet(config.Key).(config.AppConfigInterface)
	loggerService := c.MustGet(logger.Key).(*logrus.Logger).WithField("service", Key)
	c.Set(Key, func(c core.Container) interface{} {
		return NewMysqlConn(appConfig.GetMasterConfig(), appConfig.GetSlaveConfig(), loggerService)
	})
}
