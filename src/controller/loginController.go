package controller

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
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

	//对密码进行加密.
	//返回一个新的哈希散列来计算md5.
	encryption := md5.New()
	encryption.Write([]byte(password))
	password = hex.EncodeToString(encryption.Sum(nil))

	//登录逻辑.
	rs := model.Login(account, password)
	fmt.Println(rs)
}
