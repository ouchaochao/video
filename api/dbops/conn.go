package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err error
)

func init()  {
	// 复用dbConn
	dbConn, err = sql.Open("mysql", "root:1@tcp(47.94.131.35:3306)/video?charset=utf8")
	if err != nil{
		// 无法连接时抛出异常
		panic(err.Error())
	}
}
