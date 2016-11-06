package model

import (
	"fmt"
)

func Login(account string, password string) (bool, error) {
	fmt.Println("登录")
	fmt.Println("account:", account)
	fmt.Println("password:", password)

	//定义一个登录的状态.
	//这里定义true，表示登录成功.
	status := true

	//定义一个错误信息.
	msg := "登录失败"

	//验证登录的帐号.
	ac := CheckAccount(account)

	//如果验证用户名的方法传入过来的错误信息不为空，就表示用户名存在错误.
	//就需要返回错误状态.
	if ac != nil {
		status = false
	}

	//验证登录的密码.
	pw := CheckPassword(password)

	//验证登录密码是否返回的有错误信息.
	if pw != nil {
		status = false
	}

	//验证帐号密码的有效性.

	return status, msg

	//根据account来查询密码然后对比加密后的密码与传入过来的密码是否相同.

	//判断帐号和密码是否为空.

}

//判断帐号的长度是否为空等等.
//由于登录是可以使用用户名，邮箱，和手机号，所以不能验证用户名的长度，因为用户名可能为一个字符.
func CheckAccount(account string) error {

	err := "用户名不能为空"

	if len(account) != 0 {
		err = nil
	}
	return err
}

//判断登录的密码长度等等.
func CheckPassword(password string) bool {

	err := "密码长度不够"

	if len(password) > 6 {
		err = nil
	}
	return err
}
