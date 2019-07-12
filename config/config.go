package config

import (
	"encoding/json"
	"os"
)

//服务端配置
type AppConfig struct {
	//应用名称
	AppName    string `json:"app_name"`
	Port       string `json:"port"`
	StaticPath string `json:"static_path"`
	Mode       string `json:"mode"`
	SetLevel   string `json:"debug"`
	//数据库驱动
	DriverName string `json:"mysql"`
	//数据库连接
	DataSourceName string `json:"root:willzhang@/goCRM?charset=utf8"`
	//是否显示SQL语句
	ShowSQL bool `json:ture`
}

var ServConfig AppConfig

//初始化服务器配置
func InitConfig() *AppConfig {

	file, err := os.Open("./config/config.json")
	if err != nil {
		panic(err.Error())
	}
	decoder := json.NewDecoder(file)
	conf := AppConfig{}
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err.Error())
	}
	return &conf
}
