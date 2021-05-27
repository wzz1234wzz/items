package sqldao

import (
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"sqlserver/models"
	"testing"
)

//type User struct {
//	Userid int64  `gorm:"column:userid"`
//	Name   string `gorm:"column:name"`
//	Age    int64  `gorm:"column:age"`
//}

type User struct {
	Userid int64
	Name   string
	Age    int64
}

func TestGorm(t *testing.T) {
	//连接字符串
	dsn := "sqlserver://sa:123456@localhost:1433?database=wzz"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//var user models.User
	var user User
	var us models.Users
	var users []User
	var userid int64
	db.Table("dbo.users").Where("userid>?", 2).Order("userid asc").Find(&users)
	db.Table("dbo.users").Select("userid").Order("userid asc").Last(&userid)
	db.Table("dbo.users").Last(&user)
	db.Table("dbo.users").Last(&us)
	for _, u := range users {
		fmt.Println(u)
	}
	fmt.Println("user=", user)
	fmt.Println("us=", us)
	fmt.Println("userid=", userid)
}
