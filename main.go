package main

import (
	"fmt"
	"go-web/internal/config"
	"go-web/internal/db"
	"go-web/internal/http"
	"go-web/internal/logger"
)

func main() {
	appName := config.AppConfig.GetString("app.name")
	mode := config.AppConfig.GetString("app.mode")
	port := fmt.Sprintf(":%v", config.AppConfig.GetInt("app.server.port"))
	logger.Info("(启动中) 应用启动: 应用名: %s, 端口: %s, 环境:%s", appName, port, mode)
	// 初始化数据库
	db.InitDB()
	// 初始化网络服务
	app := http.NewServer(mode)
	// 启动服务器
	err := app.Run(port)
	if err != nil {
		logger.Error("%s", err)
	}
}
