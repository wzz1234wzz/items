package sqldao

import (
	"database/sql"
	"fmt"
	_ "github.com/alexbrainman/odbc"
	_ "github.com/denisenkom/go-mssqldb"
	"testing"
	"time"
)

type AccessRegion struct {
	userid int64
	name   string
	age    int64
}

func TestMssqlStruct(t *testing.T) {
	var server = "localhost"
	var port = 1433
	var user = "sa"
	var password = "123456"
	var database = "wzz"

	//连接字符串
	connString := fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s", server, port, database, user, password)

	//建立连接
	db, err := sql.Open("mssql", connString)
	if err != nil {
		t.Fatal("Open Connection failed:", err.Error())
	}
	defer db.Close()

	//通过连接对象执行查询
	rows, err := db.Query(`select * from dbo.users`)
	if err != nil {
		t.Fatal("Query failed:", err.Error())
	}
	defer rows.Close()

	var rowsData []*AccessRegion
	//遍历每一行
	for rows.Next() {
		var row = new(AccessRegion)
		rows.Scan(&row.userid, &row.name, &row.age)
		rowsData = append(rowsData, row)
	}

	//打印数组
	for _, ar := range rowsData {
		fmt.Print(ar.userid, "\t", ar.name, "\t", ar.age)
		fmt.Println()
	}
}

func TestConnect(t *testing.T) {
	var isdebug = true
	var server = "localhost"
	var port = 1433
	var user = "sa"
	var password = "123456"
	var database = "wzz"

	//连接字符串
	connString := fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s", server, port, database, user, password)
	if isdebug {
		fmt.Println(connString)
	}
	//建立连接
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		t.Fatal("Open Connection failed:", err.Error())
	}
	defer conn.Close()
	t.Log("连接成功！")
	//产生查询语句的Statement
	stmt, err := conn.Prepare(`select * from dbo.users`)
	if err != nil {
		t.Fatal("Prepare failed:", err.Error())
	}
	rows, err := stmt.Query()
	if err != nil {
		t.Fatal("Query failed:", err.Error())
	}
	defer stmt.Close()

	//建立一个列数组
	cols, err := rows.Columns()
	var colsdata = make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		colsdata[i] = new(interface{})
		fmt.Print(cols[i])
		fmt.Print("\t")
	}
	fmt.Println()

	//遍历每一行
	for rows.Next() {
		rows.Scan(colsdata...) //将查到的数据写入到这行中
		PrintRow(colsdata)     //打印此行
	}
	defer rows.Close()
}

//打印一行记录，传入一个行的所有列信息
func PrintRow(colsdata []interface{}) {
	for _, val := range colsdata {
		switch v := (*(val.(*interface{}))).(type) {
		case nil:
			fmt.Print("NULL")
		case bool:
			if v {
				fmt.Print("True")
			} else {
				fmt.Print("False")
			}
		case []byte:
			fmt.Print(string(v))
		case time.Time:
			fmt.Print(v.Format("2016-01-02 15:05:05.999"))
		default:
			fmt.Print(v)
		}
		fmt.Print("\t")
	}
	fmt.Println()
}

func TestODBC(t *testing.T) {
	db, err := sql.Open("odbc", "driver={sql server};server=localhost;port=1433;uid=sa;pwd=123456;database=wzz")
	if err != nil {
		fmt.Printf(err.Error())
	}
	//通过连接对象执行查询
	rows, err := db.Query(`select * from dbo.users`)
	if err != nil {
		t.Fatal("Query failed:", err.Error())
	}
	defer rows.Close()

	var rowsData []*AccessRegion
	//遍历每一行
	for rows.Next() {
		var row = new(AccessRegion)
		rows.Scan(&row.userid, &row.name, &row.age)
		rowsData = append(rowsData, row)
	}

	//打印数组
	for _, ar := range rowsData {
		fmt.Print(ar.userid, "\t", ar.name, "\t", ar.age)
		fmt.Println()
	}
}
