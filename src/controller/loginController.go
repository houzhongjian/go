package controller

import (
	"html/template"
	"model"
	"net/http"
	"strings"
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
	account := r.FormValue("account")
	password := r.FormValue("password")

	//去掉前后的空格.
	account = strings.Trim(account, " ")
	password = strings.Trim(password, " ")

	//登录逻辑.
	model.Login(account, password)

}
