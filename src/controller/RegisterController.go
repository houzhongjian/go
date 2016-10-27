package controller

import (
	"encoding/json"
	"html/template"
	"model"
	"net/http"
)

type RegisterController struct {
}

const tempPath = "../template/html/register/"

func (this *RegisterController) IndexAction(resp http.ResponseWriter, req *http.Request) {

	//加载视图.
	t, _ := template.ParseFiles(tempPath + "index.html")
	t.Execute(resp, nil)
}

func (this *RegisterController) RegisterAction(resp http.ResponseWriter, req *http.Request) {

	//定义一个结构体.
	type Users struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		RePassword string `json:"rePassword"`
		Status     bool   `json:"status"`
	}

	//获取访问的方式.
	method := req.Method

	if method == "POST" {
		//获取发送过来的数据.
		UserName := req.FormValue("username")
		PassWord := req.FormValue("password")
		RePassword := req.FormValue("rePassword")

		//定义一个map.
		UserData := make(map[string]string)

		//给map赋值.
		UserData["username"] = UserName
		UserData["password"] = PassWord
		UserData["rePassword"] = RePassword

		//将map传给模型来验证数据的合法性.
		sta, err := model.AddUser(UserData)

		//给结构体赋值.
		data := &Users{Username: err["username"], Password: err["password"], RePassword: err["rePassword"], Status: sta}

		//将结构体转换成json.
		res, _ := json.Marshal(data)

		//返回给请求.
		resp.Write(res)

	} else {
		http.Redirect(resp, req, "/404", http.StatusMovedPermanently)
	}

}
