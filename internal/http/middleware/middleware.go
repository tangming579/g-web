package middleware

import (
	"github.com/gin-gonic/gin"
	"go-web/internal/logger"
)

// LoadMiddleWare 加载中间件
func LoadMiddleWare(server *gin.Engine) {
	server.Use(Options, GinLogger(logger.Logger), PanicResponse())
}
