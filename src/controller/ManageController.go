package controller

import (
	//	"html/template"
	"net/http"
)

func init() {
	//验证是否有登录.
}

//结构体.
type ManageController struct {
}

//index方法.
func (this *ManageController) IndexAction(resp http.ResponseWriter, req *http.Request) {
	RequireHtml(resp, nil, "../template/html/login/login.html")
}

//func RequireHtml(resp http.ResponseWriter, data map[string]interface{}, tmpl ...string) {
//	t, _ := template.ParseFiles(tmpl)
//	t.Execute(resp, data)
//}
