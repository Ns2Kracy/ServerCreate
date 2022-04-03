package config

import (
	"encoding/json"
	"os"
)

//AppConfig 服务端设置
type AppConfig struct {
	AppName    string `json:"appName"`
	Port       string `json:"port"`
	StaticPath string `json:"staticPath"`
	Mode       string `json:"mode"`
	MySQL      MySQL  `json:"mysql"`
	Redis      Redis  `json:"redis"`
}

//MySQL 设置
type MySQL struct {
	Drive    string `json:"drive"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Host     string `json:"host"`
	Database string `json:"database"`
}

//Redis 设置
type Redis struct {
	NetWork string `json:"netWork"`
	Addr    string `json:"addr"`
	Port    string `json:"port"`
	Pwd     string `json:"pwd"`
	Prefix  string `json:"prefix"`
}

// InitConfig 初始化服务器配置
func InitConfig() *AppConfig {
	var config *AppConfig
	file, err := os.Open("./config.json")
	if err != nil {
		panic(err.Error())
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err.Error())
	}
	return config
}
