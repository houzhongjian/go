package controller

import (
	"fmt"
	"html/template"
	"model"
	"net/http"
)

type IndexController struct {
}

//定义路径常量.
const IndexPath = "../template/html/index/"
const Layout = "../template/html/layout/"

//首页方法.
func (this *IndexController) IndexAction(resp http.ResponseWriter, req *http.Request) {

	//查询分类.
	rows := model.GetCategory()

	for rows.Next() {
		var id int
		var category_name string
		var is_delete int
		var order int
		var url string
		err := rows.Scan(&id, &category_name, &is_delete, &order, &url)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(id, category_name, is_delete, order, url)
		}
	}

	t, _ := template.ParseFiles(Layout+"index.html", IndexPath+"right.html")
	value := map[string]string{"username": "侯忠建"}
	t.Execute(resp, value)
}
