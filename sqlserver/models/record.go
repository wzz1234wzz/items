package models

import (
	"time"
)

type Record struct {
	Id        int       `xorm:"not null pk autoincr INT(10)"`
	TableName string    `xorm:"not null VARCHAR(100)"`
	LastId    int64     `xorm:"not null BIGINT(20)"`
	CreatedAt time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdatedAt time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP"`
}
