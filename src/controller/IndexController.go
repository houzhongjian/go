package controller

import (
	"html/template"
	"net/http"
)

type IndexController struct {
}

//定义路径常量.
const IndexPath = "../template/html/index/"

//首页方法.
func (this *IndexController) IndexAction(resp http.ResponseWriter, req *http.Request) {

	t, _ := template.ParseFiles(IndexPath + "index.html")
	t.Execute(resp, nil)
}
