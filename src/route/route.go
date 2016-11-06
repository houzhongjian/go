package route

import (
	"controller"
	"fmt"
	"html/template"
	"net/http"
	"reflect"
	"strings"
)

//处理登录请求.
func LoginHandle(resp http.ResponseWriter, req *http.Request) {

	//获取结构体.
	login := &controller.LoginController{}

	loginController := reflect.ValueOf(login)

	//接收控制器.
	//并将控制器的首字母转换成大写.
	action := req.FormValue("a")

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

//默认访问路由.
func DefaultHandle(resp http.ResponseWriter, req *http.Request) {

	url := req.URL.Path

	if url != "/" {
		http.Redirect(resp, req, "/404", http.StatusMovedPermanently)
	} else {
		index := &controller.IndexController{}
		IndexController := reflect.ValueOf(index)

		respVal := reflect.ValueOf(resp)
		reqVal := reflect.ValueOf(req)

		//反射.
		IndexController.MethodByName("IndexAction").Call([]reflect.Value{respVal, reqVal})
	}
}

//处理404页面请求 .
func HttpNotFundHandle(resp http.ResponseWriter, req *http.Request) {
	url := req.URL.Path

	if url != "/" {
		t, _ := template.ParseFiles("../template/html/404.html")
		t.Execute(resp, nil)
	}
}

//设置注册路由.
func RegisterHandle(resp http.ResponseWriter, req *http.Request) {

	//获取url.
	url := req.URL.Path

	//去掉url前后的/.
	url = strings.Trim(url, "/")

	//将url以/拆分.
	params := strings.Split(url, "/")

	var ActionName string
	if len(params) > 1 {
		//将首字母转成大写.
		ActionName = strings.Title(params[1] + "Action")

	} else {

		ActionName = "IndexAction"
	}

	register := &controller.RegisterController{}
	RegisterController := reflect.ValueOf(register)

	regResp := reflect.ValueOf(resp)
	regReq := reflect.ValueOf(req)

	//捕获异常.
	defer func() {
		if err := recover(); err != nil {
			http.Redirect(resp, req, "/404", http.StatusMovedPermanently)
		}
	}()

	RegisterController.MethodByName(ActionName).Call([]reflect.Value{regResp, regReq})
}

//后台登录路由.
func ManageHandle(resp http.ResponseWriter, req *http.Request) {

	//获取url.
	url := req.URL.Path

	//去掉url前后的/.
	url = strings.Trim(url, "/")

	//以/的方式拆成数组.
	arr := strings.Split(url, "/")

	//判断.
	//长度大于1表示有控制器以及方法.
	var ActionName string
	if len(arr) > 1 {

		//使用填写的方法.
		ActionName = arr[1] + "Action"
	} else {
		//没有方法名就使用默认的方法名.
		ActionName = "idexAction"
	}

	ActionName = strings.Title(ActionName)

	fmt.Println(ActionName)

	manage := &controller.ManageController{}

	ManageController := reflect.ValueOf(manage)

	manResp := reflect.ValueOf(resp)

	manReq := reflect.ValueOf(req)

	//捕获异常.
	defer func() {

		//抛出异常.
		if err := recover(); err != nil {
			http.Redirect(resp, req, "/404", http.StatusMovedPermanently)
		}
	}()

	ManageController.MethodByName(ActionName).Call([]reflect.Value{manResp, manReq})

}
