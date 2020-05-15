package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

var DB *sqlx.DB

func init() {
	var err error
	DB, err = sqlx.Open(`mysql`, `root:123456@tcp(127.0.0.1:3306)/go?charset=utf8&parseTime=true`)
	if err != nil {
		fmt.Println("数据库连接错误=>", err.Error())
		os.Exit(1)
	}
	err = DB.Ping()
	if err != nil {
		fmt.Println("数据库运行错误=>", err.Error())
		os.Exit(1)
	}
}
