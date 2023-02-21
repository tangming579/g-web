package middleware

import (
	"github.com/gin-gonic/gin"
	"go-web/internal/http/response"
)

func PanicResponse() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, err any) {
		switch v := err.(type) {
		case error:
			response.Failure(c, v.Error())
		default:
			response.Failure(c, v.(string))
		}
		c.Abort()
	})
}
