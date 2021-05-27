package sqldao

import (
	"context"
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"testing"
)

var SQLDSN = "sqlserver://sa:123456@loaclhost:1433?database=wzz"

func TestUsers(t *testing.T) {
	gSession, err := gorm.Open(sqlserver.Open(SQLDSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	var dao SqlUsersDao
	usr := dao.QueryLastFromSqlserver(ctx, gSession)
	fmt.Println("last id=", usr.Userid)
	users := dao.QueryByUserIDFromSqlserver(ctx, gSession, 2)
	for _, u := range users {
		fmt.Println(u)
	}
}
