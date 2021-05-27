package models

type Users struct {
	Userid int64  `xorm:"INT(11)"`
	Name   string `xorm:"CHAR(10)"`
	Age    int    `xorm:"INT(11)"`
}
