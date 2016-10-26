package model

import (
	"db"
)

func AddUser(data map[string]string) bool {

	//获取传递进来的map.
	//将map里面的数据转换成sql.
	strSql := "INSERT INTO `user`(`id`, `username`, `password`) VALUES (NULL,'" + data["username"] + "',md5(" + data["password"] + "))"
	return db.Query(strSql)

}
