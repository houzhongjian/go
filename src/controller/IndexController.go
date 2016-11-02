package controller

import (
	//	"fmt"
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
	model.GetCategory()

	//	for rows.Next() {
	//		var category_name string
	//		var url string
	//		fmt.Println(category_name)
	//		fmt.Println(url)
	//	}

	t, _ := template.ParseFiles(Layout+"index.html", IndexPath+"right.html")
	value := map[string]string{"username": "侯忠建"}
	t.Execute(resp, value)
}
