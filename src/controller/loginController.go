package controller

import (
	"fmt"
	"html/template"
	"net/http"
)

type LoginController struct {
}

func (this *LoginController) IndexAction(w http.ResponseWriter, r *http.Request) {

	//加载视图文件.
	t, _ := template.ParseFiles("../template/html/login/login.html")
	t.Execute(w, nil)
}

func (this *LoginController) LoginAction(w http.ResponseWriter, r *http.Request) {
	//接收发送过来的用户名和密码.
	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Println(username, ":", password)
}
