package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	configPath = kingpin.Flag("config", "path to config file").Default(".").String()
	AppConfig  *viper.Viper
)

func init() {
	loadAppConfig()
}

func loadAppConfig() {
	AppConfig = viper.New()
	AppConfig.SetConfigName("app")
	AppConfig.SetConfigType("yaml")
	AppConfig.AddConfigPath(*configPath)
	err := AppConfig.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("读取配置文件失败:%s", err))
	}
	AppConfig.WatchConfig()
}
