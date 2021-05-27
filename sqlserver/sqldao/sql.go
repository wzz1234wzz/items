package sqldao

import (
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"sqlserver/config"
)

var Sql *gorm.DB
var SqlUser = &SqlUsersDao{}

func InitSqlserver() {
	gSession, err := gorm.Open(sqlserver.Open(config.Custom.SqlServer.DNS), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Sql = gSession
	fmt.Println("InitSqlserver success!")
}
