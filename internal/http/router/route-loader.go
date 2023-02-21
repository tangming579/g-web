package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/internal/logger"
	"net/http"
)

var RouteInfo = RouteMappings{
	routes: make(map[string]func(*gin.RouterGroup)),
}

type RouteMappings struct {
	routes map[string]func(*gin.RouterGroup)
}

func (mappings *RouteMappings) addRoute(groupPath string, fn func(*gin.RouterGroup)) {
	if nil != mappings.routes[groupPath] {
		panic(fmt.Sprintf("定义了重复的路由分组:%s", groupPath))
	}
	mappings.routes[groupPath] = fn
}

func (mappings *RouteMappings) getRoutes() map[string]func(*gin.RouterGroup) {
	return mappings.routes
}

// LoadRoutes 加载路由
func LoadRoutes(server *gin.Engine) {
	// 404
	server.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "没有找到对应的路由")
	})

	for key, info := range RouteInfo.getRoutes() {
		logger.Info("(启动中) 加载路由表:%s", key)
		group := server.Group(key)
		info(group)
	}
}
