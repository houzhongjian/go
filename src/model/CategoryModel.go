package model

import (
	"database/sql"
	"db"
	//	"fmt"
)

//未删除.
const DELETE = "0"

//已删除.
const NOT_DELETE = "1"

func GetCategory() *sql.Rows {

	//定义查询首页类别的sql.
	sql := "SELECT `Id`,`Category_name`,`Is_delete`,`Order`,`Url` FROM `category` WHERE `Is_delete` = " + DELETE + " ORDER BY `Order` DESC"

	//获取db对象.
	dbObj := db.GetInstance()

	rows, _ := dbObj.Query(sql)

	return rows

	//	fmt.Println("sql:", sql)
	//	for rows.Next() {
	//		var id int
	//		var category_name string
	//		var is_delete int
	//		var order int
	//		var url string

	//		err := rows.Scan(&id, &category_name, &is_delete, &order, &url)

	//		if err != nil {
	//			fmt.Println("err:", err)
	//		}

	//		fmt.Println(id, category_name, is_delete, order, url)
	//	}
}
