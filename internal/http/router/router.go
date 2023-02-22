package router

import (
	"github.com/gin-gonic/gin"
	"go-web/internal/controller"
)

func init() {
	RouteInfo.addRoute("api", func(group *gin.RouterGroup) {
		group.GET("/user/", controller.GetUserInfo)
	})
}
