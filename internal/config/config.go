package config

const (
	fieldConfigError = "empty %s value in database.%s.%s in config file"
)

type AppConfigInterface interface {
	GetMasterConfig() *DatabaseConfig
	GetSlaveConfig() *DatabaseConfig
}

func New(masterConfig DatabaseConfig, slaveConfig DatabaseConfig) *appConfig {
	return &appConfig{
		master: masterConfig,
		slave:  slaveConfig,
	}
}

type appConfig struct {
	master DatabaseConfig
	slave  DatabaseConfig
}

type DatabaseConfig struct {
	Addr      string `yaml:"addr"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	DbName    string `yaml:"dbname"`
	Charset   string `yaml:"charset"`
	Collation string `yaml:"collation"`
	Timeout   int64  `yaml:"timeout"`
}

func (a *appConfig) GetMasterConfig() *DatabaseConfig {
	return &a.master
}

func (a *appConfig) GetSlaveConfig() *DatabaseConfig {
	return &a.slave
}
