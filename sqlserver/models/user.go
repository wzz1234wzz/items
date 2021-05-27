package models

import (
	"time"
)

type User struct {
	Ysid        int64     `xorm:"not null pk autoincr BIGINT(20)"`
	Name        string    `xorm:"not null unique VARCHAR(128)"`
	PwdHash     string    `xorm:"VARCHAR(128)"`
	Group       string    `xorm:"not null VARCHAR(20)"`
	Rank        int       `xorm:"default 1 TINYINT(1)"`
	Email       string    `xorm:"VARCHAR(128)"`
	Phone       string    `xorm:"VARCHAR(16)"`
	CreatedAt   time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdatedAt   time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP"`
	IsActivated int       `xorm:"default 1 TINYINT(1)"`
	IsAdmin     int       `xorm:"default 0 TINYINT(1)"`
}
