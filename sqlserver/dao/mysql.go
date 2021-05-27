package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sqlserver/config"
	"xorm.io/xorm"
)

var Orm *xorm.Engine

var (
	Users  = &UsersDao{}
	Record = &RecordDao{}
)

func InitMysql() {
	orm, err := xorm.NewEngine("mysql", config.Custom.Mysql.DNS)
	if err != nil {
		fmt.Println("init mysql fail. err = ", err)
		return
	}
	Orm = orm
	fmt.Println("InitMysql success!")
}
