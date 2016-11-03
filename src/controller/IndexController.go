package controller

import (
	//	"database/sql"
	//	"fmt"
	"html/template"
	"model"
	"net/http"
)

type Data struct {
	Id            int
	Category_name string
	Is_delete     int
	Order         int
	Url           string
}
type IndexController struct {
}

//定义路径常量.
const IndexPath = "../template/html/index/"
const Layout = "../template/html/layout/"

//首页方法.
func (this *IndexController) IndexAction(resp http.ResponseWriter, req *http.Request) {

	//查询分类.
	rows := model.GetCategory()

	t, _ := template.ParseFiles(Layout+"index.html", IndexPath+"right.html")
	t.Execute(resp, map[string]interface{}{"Rows": rows})
}
