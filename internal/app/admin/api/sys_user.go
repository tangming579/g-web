package api

import (
	"github.com/gin-gonic/gin"
	"go-web/internal/http/response"
)

func GetUserInfo(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		response.Failure(c, "userId不能为空")
		return
	}
}
