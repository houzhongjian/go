package controller

import (
	"crypto/md5"
	"fmt"
	"html/template"
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
	username := r.FormValue("username")
	password := r.FormValue("password")

	//去掉前后的空格.
	username = strings.Trim(username, " ")
	password = strings.Trim(password, " ")

	//对密码进行加密.
	encryption := md5.New()
	fmt.Println(string(encryption.Sum([]byte(password))))

}
