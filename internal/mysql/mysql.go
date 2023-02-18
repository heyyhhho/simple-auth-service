package mysql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"time"
)

const (
	charsetDefaultInfo   = "empty charset value in in config file, using default value %s"
	collationDefaultInfo = "empty collation value in config file, using default value %s"
	timeoutDefaultInfo   = "empty timeout value in config file, using default value %s"
	defaultCharset       = "utf8mb4"
	defaultCollation     = "utf8mb4_unicode_ci"
	defaultTimeout       = time.Second
)

type Connections interface {
	GetMasterConn() Adapter
	GetSlaveConn() Adapter
}

type Adapter interface {
	Begin() (*sql.Tx, error)
	BeginTx(context.Context, *sql.TxOptions) (*sql.Tx, error)
	Close() error
	Conn(context.Context) (*sql.Conn, error)
	Driver() driver.Driver
	Exec(string, ...interface{}) (sql.Result, error)
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	Ping() error
	PingContext(context.Context) error
	Prepare(string) (*sql.Stmt, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
	SetConnMaxLifetime(time.Duration)
	SetMaxIdleConns(int)
	SetMaxOpenConns(int)
	Stats() sql.DBStats
}

type sqlAdapter struct {
	master Adapter
	slave  Adapter
}

func (c *sqlAdapter) GetMasterConn() Adapter {
	return c.master
}

func (c *sqlAdapter) GetSlaveConn() Adapter {
	return c.slave
}
