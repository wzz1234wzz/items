package dao

import (
	"context"
	"errors"
	"sqlserver/models"
	"xorm.io/xorm"
)

type UsersDao struct{}

func (dao *UsersDao) BatchAddUser(ctx context.Context, session *xorm.Session, users []*models.Users) error {
	if users == nil {
		return errors.New("Jobs object are nil.\n")
	}
	_, err := session.Insert(users)
	return err
}

func (dao *UsersDao) QueryByUserID(ctx context.Context, session *xorm.Session, userID int64) *models.Users {
	user := &models.Users{}
	user.Userid = userID
	has, err := session.Get(user)
	if err != nil || !has {
		return nil
	}
	return user
}
