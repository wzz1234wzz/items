package logic

import (
	"context"
	"fmt"
	cron "github.com/robfig/cron/v3"
	"sqlserver/dao"
	"sqlserver/sqldao"
)

const (
	cronSpecCheck = "*/10 * * * * *" //Every 30s
)

func Do() *cron.Cron {
	cron := cron.New(cron.WithSeconds())
	cron.AddFunc(cronSpecCheck, Supervisor)
	cron.Start()
	return cron
}

func Supervisor() {
	ctx := context.Background()
	session := dao.Orm.Context(ctx)
	record := dao.Record.QueryLastIdByTableName(ctx, session, "users")
	if record == nil {
		fmt.Println("查询失败！")
		return
	}
	fmt.Println("lastid=", record.LastId)
	users := sqldao.SqlUser.QueryByUserIDFromSqlserver(ctx, sqldao.Sql, record.LastId)
	if len(users) == 0 {
		fmt.Println("无更新...")
		return
	}
	fmt.Println("users=", users)
	for _, use := range users {
		fmt.Println(*use)
	}
	record.LastId = users[len(users)-1].Userid
	dao.Record.UpdateLastId(ctx, session, record)
	err := dao.Users.BatchAddUser(ctx, session, users)
	if err != nil {
		fmt.Println("添加失败！")
		return
	}
}
