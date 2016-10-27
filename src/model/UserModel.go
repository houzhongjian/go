package model

import (
	"db"
	//	"fmt"
)

func AddUser(data map[string]string) (bool, map[string]string) {

	//获取传递进来的map.
	//将map里面的数据转换成sql.
	errMap := make(map[string]string)

	//	fmt.Println(data)

	for k, v := range data {
		if v == "" {
			errMap[k] = UserMapFunc(k) + "不能为空"
		}
	}

	if data["password"] != data["rePassword"] {
		errMap["password"] = "两次密码不相同"
	}

	if len(errMap) > 0 {
		return false, errMap
	}

	strSql := "INSERT INTO `user`(`id`, `username`, `password`) VALUES (NULL,'" + data["username"] + "',md5('" + data["password"] + "'))"
	return db.Query(strSql), errMap
}

//map.
func UserMapFunc(str string) string {

	UserMap := make(map[string]string)

	UserMap["username"] = "用户名"
	UserMap["password"] = "密码"
	UserMap["rePassword"] = "重复密码"

	return UserMap[str]
}
