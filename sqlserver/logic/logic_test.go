package logic

import (
	"github.com/robfig/cron/v3"
	"sqlserver/config"
	"sqlserver/dao"
	"sqlserver/sqldao"
	"testing"
	"time"
)

func TestLogic(t *testing.T) {
	config.InitConfig()
	dao.InitMysql()
	sqldao.InitSqlserver()
	Do()
}

func TestCron(t *testing.T) {
	c := cron.New()
	i := 1
	c.AddFunc("*/1 * * * *", func() {
		t.Log("每分钟执行一次", i)
		i++
	})
	c.Start()
	time.Sleep(time.Minute * 5)
}
