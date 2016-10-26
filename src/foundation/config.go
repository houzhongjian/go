package foundation

import (
	"encoding/json"
	"io/ioutil"
)

//读取配置信息.
func ReadConfig() map[string]string {

	//读取配置文件.
	res, _ := ioutil.ReadFile("./config/config.json")

	//将json字符串转换成map.
	var data map[string]string

	err := json.Unmarshal(res, &data)

	if err != nil {
		return nil
	} else {
		return data
	}
}
