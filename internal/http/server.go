package http

import (
	"github.com/gin-gonic/gin"
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
