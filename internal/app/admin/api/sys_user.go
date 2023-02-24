package api

import (
	"github.com/gin-gonic/gin"
	"go-web/internal/http/response"
)

func GetUserInfo(c *gin.Context) {
	userId := c.Param("userId")
	count := c.Query("count")
	if userId == "" {
		response.Failure(c, "userId不能为空")
		return
	}
	response.Ok(c, gin.H{"hello": userId, "count": count})
	return
}
