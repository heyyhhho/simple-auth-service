package main

import (
	"fmt"
	"github.com/heyyhhho/simple-auth-service/internal/config"
	"github.com/heyyhhho/simple-auth-service/internal/core"
	"github.com/heyyhhho/simple-auth-service/internal/logger"
	"github.com/heyyhhho/simple-auth-service/internal/mysql"
)

func main() {
	app := core.NewApp()
	app.Register(&logger.Provider{})
	app.Register(&config.Provider{})
	app.Register(&mysql.Provider{})

	//test mysql conn
	err := app.MustGet(mysql.Key).(mysql.Connections).GetMasterConn().Ping()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Ping is done")
	}

	//тут потому будет bootable реализация для некоторых провайдеров. к примеру для выполнения миграций, для запуска ххтп сервиса и т д
	if err := app.Boot(); err != nil {
		panic(err)
	}
}
