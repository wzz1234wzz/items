package models

type CasbinRule struct {
	PType string `xorm:"index VARCHAR(100)"`
	V0    string `xorm:"index VARCHAR(100)"`
	V1    string `xorm:"index VARCHAR(100)"`
	V2    string `xorm:"index VARCHAR(100)"`
	V3    string `xorm:"index VARCHAR(100)"`
	V4    string `xorm:"index VARCHAR(100)"`
	V5    string `xorm:"index VARCHAR(100)"`
}
