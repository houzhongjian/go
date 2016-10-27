package controller

import (
	//	"fmt"
	//	"model"
	"net/http"
)

type IndexController struct {
}

func (this *IndexController) IndexAction(resp http.ResponseWriter, req *http.Request) {

	//	userData := make(map[string]string)

	//	userData["username"] = "张三"
	//	userData["password"] = "123"

	//	res := model.AddUser(userData)

	//	if res {
	//		resp.Write([]byte("成功"))
	//	} else {
	//		resp.Write([]byte("失败"))
	//	}
	resp.Write([]byte("首页"))
}
