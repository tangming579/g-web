package router

import (
	"github.com/gin-gonic/gin"
	"go-web/internal/app/admin/api"
)

func init() {
	RouteInfo.addRoute("api", func(group *gin.RouterGroup) {
		group.GET("/user/", api.GetUserInfo)
	})
}
