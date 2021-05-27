package config

import (
	"testing"
)

func TestUsers(t *testing.T) {
	t.Log(CurrentFileDir())
	InitConfig()
	t.Log("mysqlDNS=", Custom.Mysql.DNS)
	t.Log("sqlserverDNS=", Custom.SqlServer.DNS)
}
