package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewConfig(logger logrus.FieldLogger) AppConfigInterface {

	dirName := "configs/"
	fileName := "conf"
	logger.Debugf("reading config from dir: %s, file: %s", dirName, fileName)
	viper.SetConfigName(fileName)
	viper.AddConfigPath(dirName)
	err := viper.ReadInConfig()

	if err != nil {
		logger.Fatalf("error while reading conf file %s in dir %s", fileName, dirName)
	}

	masterDatabaseConfig := DatabaseConfig{}
	slaveDatabaseConfig := DatabaseConfig{}

	err = viper.UnmarshalKey("database.master", &masterDatabaseConfig)

	if err != nil {
		if err != nil {
			logger.Fatalln("error while unmarshal key database.master")
		}
	}

	if masterDatabaseConfig.Addr == "" {
		logger.Fatalf(fieldConfigError, "addr", "master", "addr")
	}

	if masterDatabaseConfig.DbName == "" {
		logger.Fatalf(fieldConfigError, "dbname", "master", "dbname")
	}

	if masterDatabaseConfig.Username == "" {
		logger.Fatalf(fieldConfigError, "username", "master", "username")
	}

	if masterDatabaseConfig.Password == "" {
		logger.Fatalf(fieldConfigError, "password", "master", "password")
	}

	if viper.IsSet("database.master.timeout") {
		masterDatabaseConfig.Timeout = viper.GetInt64("database.master.timeout")
	}

	err = viper.UnmarshalKey("database.slave", &slaveDatabaseConfig)

	if err != nil {
		if err != nil {
			logger.Fatalln("error while unmarshal key database.slave")
		}
	}

	if slaveDatabaseConfig.Addr == "" {
		logger.Fatalf(fieldConfigError, "addr", "slave", "addr")
	}

	if slaveDatabaseConfig.DbName == "" {
		logger.Fatalf(fieldConfigError, "dbname", "slave", "dbname")
	}

	if slaveDatabaseConfig.Username == "" {
		logger.Fatalf(fieldConfigError, "username", "slave", "username")
	}

	if slaveDatabaseConfig.Password == "" {
		logger.Fatalf(fieldConfigError, "password", "slave", "password")
	}

	if viper.IsSet("database.slave.timeout") {
		masterDatabaseConfig.Timeout = viper.GetInt64("database.slave.timeout")
	}

	return New(masterDatabaseConfig, slaveDatabaseConfig)
}
