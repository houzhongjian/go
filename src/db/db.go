package db

import (
	"database/sql"
	"foundation"

	_ "github.com/go-sql-driver/mysql"
)

//定义一个变量用来存储数据库的连接资源.
var dbInstance *sql.DB

//执行sql语句的方法.
func Query(sqlStr string) bool {

	//调用单例模式的方法.
	db := GetInstance()
	//	sqlStr := "INSERT INTO `user`(`id`,`username`,`password`)VALUES(null, '张三', md5(123))"
	_, err := db.Exec(sqlStr)

	var queryStatus bool = true

	if err != nil {
		queryStatus = false
	}

	return queryStatus

}

//连接数据库的单例模式.
func GetInstance() *sql.DB {

	if dbInstance == nil {

		//获取配置信息.
		//data为map类型.
		data := foundation.ReadConfig()
		db, _ := sql.Open(data["type"], data["username"]+":"+data["password"]+"@tcp("+data["host"]+":"+data["port"]+")/"+data["dbName"]+"?charset="+data["charset"])

		dbInstance = db
	}
	return dbInstance
}
