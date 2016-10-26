package main

import (
	"fmt"
	"model"
)

type User struct {
	id       int
	username string
	password string
}

func main() {

	s := make(map[string]string)
	s["username"] = "张三"
	s["password"] = "123"
	
	
	sta := model.AddUser(s map[string]string)

}
