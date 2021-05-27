package main

import (
	"sqlserver/config"
	"sqlserver/dao"
	"sqlserver/logic"
	"sqlserver/sqldao"
)

func init() {
	config.InitConfig()
	dao.InitMysql()
	sqldao.InitSqlserver()
}

func main() {
	cron := logic.Do()
	defer cron.Stop()

	// 死循环
	ch := make(chan int)
	<-ch
}
