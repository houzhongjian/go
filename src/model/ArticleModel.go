package model

import (
	"database/sql"
	"db"
)

const ARTICLE_DELETE = "1"
const ARTICLE_NOT_DELETE = "0"

func GetArticle() *sql.Rows {

	obj := db.GetInstance()

	sql := "SELECT * FROM `article` WHERE `is_delete` = " + ARTICLE_NOT_DELETE + " LIMIT 10"

	rows, _ := obj.Query(sql)

	return rows
}
