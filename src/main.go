package main

import (
	"net/http"
	"route"
)

func main() {

	//设置登录路由.
	http.HandleFunc("/login", route.LoginHandle)

	//设置默认访问路由.
	http.HandleFunc("/", route.DefaultHandle)

	//设置404页面的路由.
	http.HandleFunc("/404", route.HttpNotFundHandle)

	//设置注册路由.
	http.HandleFunc("/register/", route.RegisterHandle)

	//设置后台登录路由.
	http.HandleFunc("/manage/", route.ManageHandle)

	//设置访问详细的目录.
	http.HandleFunc("/detail/", route.DetailHandle)

	http.Handle("/js/", http.FileServer(http.Dir("../template")))
	http.Handle("/images/", http.FileServer(http.Dir("../template")))
	http.Handle("/css/", http.FileServer(http.Dir("../template")))

	//监听80端口.
	http.ListenAndServe(":80", nil)

}
