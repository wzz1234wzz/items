package sqldao

import (
	"context"
	"gorm.io/gorm"
	"sqlserver/models"
)

type SqlUsersDao struct{}

func (dao *SqlUsersDao) QueryByUserIDFromSqlserver(ctx context.Context, gSession *gorm.DB, lastId int64) []*models.Users {
	var users = make([]*models.Users, 0)
	gSession.Table("dbo.users").Where("userid> ?", lastId).Find(&users)
	return users
}

func (dao *SqlUsersDao) QueryLastFromSqlserver(ctx context.Context, gSession *gorm.DB) models.Users {
	var user models.Users
	gSession.Table("dbo.users").Last(&user)
	return user
}
