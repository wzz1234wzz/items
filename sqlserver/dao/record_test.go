package dao

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"sqlserver/models"
	"testing"
	"time"
	"xorm.io/xorm"
)

var mysqlDNS = "root:123456@tcp(10.0.0.0:3306)/wzz_db?charset=utf8"

func TestAddRecord(t *testing.T) {
	ctx := context.Background()
	orm, err := xorm.NewEngine("mysql", mysqlDNS)
	if err != nil {
		t.Fatal("open mysql fail. err = ", err)
	}
	session := orm.Context(ctx)
	var dao RecordDao
	err = dao.AddTable(ctx, session, "job", 0)
	if err != nil {
		t.Fatal("AddTable fail. err = ", err)
	}
}

func TestQueryLastIdByTableName(t *testing.T) {
	ctx := context.Background()
	orm, err := xorm.NewEngine("mysql", mysqlDNS)
	if err != nil {
		t.Fatal("open mysql fail. err = ", err)
	}
	session := orm.Context(ctx)
	var dao RecordDao
	record := dao.QueryLastIdByTableName(ctx, session, "users")
	if record == nil {
		t.Fatal("QueryLastIdByTableName fail. err = ", err)
	}
	t.Log("lastId=", record.LastId)
}

func TestUpdateLastId(t *testing.T) {
	ctx := context.Background()
	orm, err := xorm.NewEngine("mysql", mysqlDNS)
	if err != nil {
		t.Fatal("open mysql fail. err = ", err)
	}
	session := orm.Context(ctx)
	var dao RecordDao
	record := &models.Record{1, "users", 5, time.Now(), time.Now()}
	err = dao.UpdateLastId(ctx, session, record)
	if err != nil {
		t.Fatal("UpdateLastId fail. err = ", err)
	}
	record = dao.QueryLastIdByTableName(ctx, session, "users")
	if record == nil {
		t.Fatal("QueryLastIdByTableName fail. err = ", err)
	}
	t.Log("lastId=", record.LastId)
}
