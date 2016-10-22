package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

func main() {

	//设置登录路由.
	http.HandleFunc("/login", LoginHandle)

	//设置默认访问路由.
	http.HandleFunc("/", DefaultHandle)

	//设置404页面的路由.
	http.HandleFunc("/404", HttpNotFundHandle)

	//监听80端口.
	http.ListenAndServe(":80", nil)

}

//处理登录请求.
func LoginHandle(resp http.ResponseWriter, req *http.Request) {

	//获取结构体.
	login := &loginController{}
	loginController := reflect.ValueOf(login)

	//接收控制器.
	//并将控制器的首字母转换成大写.
	action := req.FormValue("a")
	fmt.Println(action)

	if action == "" {
		action = "index"
	}

	actionName := strings.Title(action) + "Action"

	//获取数据转换格式.
	respVal := reflect.ValueOf(resp)
	reqVal := reflect.ValueOf(req)

	//捕获异常.
	//跳转到404.
	defer func() {
		if err := recover(); err != nil {
			http.Redirect(resp, req, "/404", http.StatusMovedPermanently)
		}
	}()

	//反射调用方法，传递数据到调用的方法.
	//如果传入的方法名不存在就会抛出一个异常.
	loginController.MethodByName(actionName).Call([]reflect.Value{respVal, reqVal})

}

type loginController struct {
}

func (this *loginController) IndexAction(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("显示登录页面"))
}

func (this *loginController) LoginAction(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("处理登录业务请求"))
}

func DefaultHandle(resp http.ResponseWriter, req *http.Request) {

	url := req.URL.Path
	fmt.Println(url)

	if url != "/" {
		http.Redirect(resp, req, "/404", http.StatusMovedPermanently)
	} else {
		resp.Write([]byte("首页"))
	}
}

//处理404页面请求 .
func HttpNotFundHandle(resp http.ResponseWriter, req *http.Request) {
	url := req.URL.Path

	if url != "/" {
		resp.Write([]byte("<h1 align='center'>404</h1>"))
	}
}
