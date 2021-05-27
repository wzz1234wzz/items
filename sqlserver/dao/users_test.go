package dao

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"sqlserver/models"
	"testing"
	"xorm.io/xorm"
)

func TestBatchAddUsers(t *testing.T) {
	var users []*models.Users
	users = append(users, &models.Users{6, "wzz", 10})
	users = append(users, &models.Users{7, "gh", 20})

	ctx := context.Background()
	orm, err := xorm.NewEngine("mysql", mysqlDNS)
	if err != nil {
		t.Fatal("open mysql fail. err = ", err)
	}
	session := orm.Context(ctx)
	var dao UsersDao
	err = dao.BatchAddUser(ctx, session, users)
	if err != nil {
		t.Fatal("AddTable fail. err = ", err)
	}
	user := dao.QueryByUserID(ctx, session, 7)
	t.Log("user=", user)
}
