package router

import (
	"github.com/gin-gonic/gin"
	"go-web/internal/app/admin/api"
)

func init() {
	RouteInfo.addRoute("/api/user", func(group *gin.RouterGroup) {
		group.GET("/:userId", api.GetUserInfo)
		group.GET("", api.GetUserList)
		group.DELETE("/:userId", api.DeleteUser)
		group.PUT("/:userId", api.UpdateUser)
		group.POST("", api.CreateUser)
	})
}
