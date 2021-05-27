package dao

import (
	"context"
	"sqlserver/models"
	"time"
	"xorm.io/xorm"
)

type RecordDao struct{}

func (dao *RecordDao) AddTable(ctx context.Context, session *xorm.Session, tableName string, lastId int64) error {
	record := &models.Record{}
	record.TableName = tableName
	record.LastId = lastId
	record.CreatedAt = time.Now()
	record.UpdatedAt = time.Now()
	_, err := session.InsertOne(record)
	return err
}

func (dao *RecordDao) QueryLastIdByTableName(ctx context.Context, session *xorm.Session, tableName string) *models.Record {
	record := &models.Record{}
	record.TableName = tableName
	has, err := session.Get(record)
	if err != nil || !has {
		return nil
	}
	return record
}

func (dao *RecordDao) UpdateLastId(ctx context.Context, session *xorm.Session, record *models.Record) error {
	session.Where("table_name = ?", record.TableName)
	_, err := session.Cols("last_id", "updated_at").Update(record)
	if err != nil {
		return err
	}
	return nil
}
