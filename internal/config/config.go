package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	configPath = *kingpin.Flag("config", "path to config file").Default("./config/app.yaml").String()
	AppConfig  *viper.Viper
)

func init() {
	loadAppConfig()
}

func loadAppConfig() {
	//初始化viper
	AppConfig = viper.New()
	//设置文件名
	AppConfig.SetConfigName("app")
	//设置文件类型
	AppConfig.SetConfigType("yaml")
	configPath = "config"
	//设置文件所在路径
	AppConfig.AddConfigPath(configPath)
	err := AppConfig.ReadInConfig()
	if err != nil {
		//手动触发宕机，让程序崩溃
		panic(fmt.Sprintf("读取配置文件失败:%s", err))
	}
	AppConfig.WatchConfig()
}
