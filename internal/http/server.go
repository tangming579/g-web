package http

import (
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"go-web/internal/config"
	"go-web/internal/http/middleware"
	"go-web/internal/http/router"
)

func NewServer(mode string) *gin.Engine {
	gin.SetMode(mode)
	server := gin.New()
	middleware.LoadMiddleWare(server)
	router.LoadRoutes(server)
	return server
}

func NewConsulClient() (*api.Client, error) {
	cCfg := &api.Config{
		Address: config.AppConfig.Get("consul.address"),
		Scheme:  config.AppConfig.Get("consul.scheme"),
		Token:   config.AppConfig.Get("consul.token"),
	}
	clt, err := api.NewClient(cCfg)
	if err != nil {
		return nil, err
	}
	return clt, nil
}
