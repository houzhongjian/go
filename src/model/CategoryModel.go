package model

import (
	"db"
	//	"fmt"
)

//未删除.
const DELETE = "0"

//已删除.
const NOT_DELETE = "1"

func GetCategory() Rows {

	//定义查询首页类别的sql.
	sql := "SELECT * FROM `category` WHERE `is_delete` = " + DELETE + " ORDER BY `order` DESC"

	//获取db对象.
	dbObj := db.GetInstance()

	Rows, _ := dbObj.Query(sql)

	return Rows

	//	fmt.Println(rows.Next())
	//	fmt.Println("SQL:", sql)
	//	fmt.Println(rows, err)
}
