package db

import (
	"fmt"
	"github.com/StupidTAO/DIDHub/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//数据库指针
var DB *sqlx.DB

func InitDB() *sqlx.DB {
	if DB != nil {
		return DB
	}

	//从配置文件中读取
	user := "root"
	password := "12345678"
	dbname := "welfare"

	_dataSourceName := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", user, password, dbname)
	// 数据源语法："用户名:密码@[连接方式](主机名:端口号)/数据库名"
	database, err := sqlx.Open("mysql", _dataSourceName)
	if err != nil {
		log.Info("open mysql failed,", err)
	}
	DB = database
	return DB
}
