package model

import (
	"db"
	//	"fmt"
	//	"fmt"
)

func AddUser(data map[string]string) (bool, map[string]string) {

	//获取传递进来的map.
	//将map里面的数据转换成sql.
	errMap := make(map[string]string)

	//验证字段是否为空.
	for k, v := range data {
		if v == "" {
			errMap[k] = UserMapFunc(k) + "不能为空"
		}
	}

	//验证两次密码是否相同.
	if data["password"] != data["rePassword"] {
		errMap["password"] = "两次密码不相同"
	}

	//验证用户名是否存在.
	errMap = ValidateUserNameUnique(data["username"], errMap)

	//验证用户名长度.
	errMap = ValidateUserName(data["username"], errMap)

	//验证密码长度.
	errMap = ValidatePassword(data["password"], errMap)

	errMap = ValidateRePassword(data["rePassword"], errMap)

	//判断是否有错误信息.
	if len(errMap) > 0 {
		return false, errMap
	}

	strSql := "INSERT INTO `user`(`id`, `username`, `password`) VALUES (NULL,'" + data["username"] + "',md5('" + data["password"] + "'))"

	//获取db对象.
	dbObj := db.GetInstance()
	_, errs := dbObj.Exec(strSql)

	status := true

	if errs != nil {
		status = false
	}
	return status, errMap
}

//验证重复密码长度.
func ValidateRePassword(rePassword string, errMap map[string]string) map[string]string {
	//验证重复密码长度.
	if len(rePassword) < 6 {
		errMap["rePassword"] = "重复密码不能小于6位数"
	}
	if len(rePassword) > 20 {
		errMap["rePassword"] = "重复密码不能大于20位"
	}
	return errMap
}

//验证密码长度.
func ValidatePassword(password string, errMap map[string]string) map[string]string {
	//验证重复密码长度.
	if len(password) < 6 {
		errMap["password"] = "重复密码不能小于6位数"
	}
	if len(password) > 20 {
		errMap["password"] = "重复密码不能大于20位"
	}
	return errMap
}

//验证用户名.
func ValidateUserName(username string, errMap map[string]string) map[string]string {

	if len(username) < 6 {
		errMap["username"] = "用户名不能小于6位数"
	}
	if len(username) > 20 {
		errMap["username"] = "用户名不能大于20位"
	}
	return errMap
}

//验证用户名是否唯一.
func ValidateUserNameUnique(username string, errMap map[string]string) map[string]string {

	//定义一个sql 用来查询.
	getUserSql := "SELECT * FROM `user` WHERE `username` = ?"

	//获取db对象.
	dbObj := db.GetInstance()

	//查询.
	row, _ := dbObj.Query(getUserSql, username)

	//验证是否存在.
	if row.Next() == true {
		errMap["username"] = "用户名已经存在"
	}

	return errMap
}

//用来翻译字段.
func UserMapFunc(str string) string {

	UserMap := make(map[string]string)

	UserMap["username"] = "用户名"
	UserMap["password"] = "密码"
	UserMap["rePassword"] = "重复密码"

	return UserMap[str]
}

//根据用户名查询用户信息.
//该方法需要传入一个字符串.
//func GetUserInfoByUserName(username string) (*sql.Row, error) {

//定义一条sql语句根据username来查询所有的用户信息.
//	sqlStr := "SELECT * FROM `user` WHERE `username` = ?"

//使用db包来调用FindOne方法来查询数据.
//FindOne方法需要传入一个sql语句.
//FindOne方法会返回两个参数一个为*sql.Row类型的结果集，一个为error类型的错误信息.
//	return db.FindOne(sqlStr)
//	db := db.GetInstance()
//	db.QueryRow()
//}
