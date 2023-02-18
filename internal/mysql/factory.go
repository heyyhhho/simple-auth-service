package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/heyyhhho/simple-auth-service/internal/config"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

func NewMysqlConn(master *config.DatabaseConfig, slave *config.DatabaseConfig, logger logrus.FieldLogger) *sqlAdapter {
	logger.Infoln("Init master conn")
	masterConn, err := sql.Open("mysql", makeDsn(master, logger))
	if err != nil {
		logger.Fatalln(err)
	}
	logger.Infoln("Init slave conn")
	slaveConn, err := sql.Open("mysql", makeDsn(slave, logger))
	if err != nil {
		logger.Fatalln(err)
	}
	return &sqlAdapter{
		master: masterConn,
		slave:  slaveConn,
	}
}

func makeDsn(databaseConfig *config.DatabaseConfig, logger logrus.FieldLogger) string {
	builder := strings.Builder{}
	builder.WriteString(databaseConfig.Username)
	builder.WriteByte(':')
	builder.WriteString(databaseConfig.Password)
	builder.WriteString("@tcp(")
	builder.WriteString(databaseConfig.Addr)
	builder.WriteString(")/")
	builder.WriteString(databaseConfig.DbName)
	if databaseConfig.Timeout == 0 {
		builder.WriteString("?timeout=")
		builder.WriteString(time.Duration(databaseConfig.Timeout).String())
	} else {
		builder.WriteString("?timeout=")
		builder.WriteString(defaultTimeout.String())
		logger.Infof(timeoutDefaultInfo, defaultTimeout)
	}

	if databaseConfig.Charset != "" {
		builder.WriteString("&charset=")
		builder.WriteString(databaseConfig.Charset)
	} else {
		builder.WriteString("&charset=")
		builder.WriteString(defaultCharset)
		logger.Infof(charsetDefaultInfo, defaultCharset)
	}
	if databaseConfig.Collation != "" {
		builder.WriteString("&collation=")
		builder.WriteString(databaseConfig.Collation)
	} else {
		builder.WriteString("&collation=")
		builder.WriteString(defaultCollation)
		logger.Infof(collationDefaultInfo, defaultCollation)
	}

	return builder.String()
}
