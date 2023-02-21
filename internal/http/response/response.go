package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiResponse struct {
	Status        int         `json:"status"`
	Message       string      `json:"message"`
	DetailMessage string      `json:"detailMessage"`
	Data          interface{} `json:"data"`
}

func Ok(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ApiResponse{
		Status: 0,
		Data:   data,
	})
}

func Failure(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, ApiResponse{
		Status:  -1,
		Message: msg,
	})
}

func FailureWithDetail(c *gin.Context, msg, detail string) {
	c.JSON(http.StatusOK, ApiResponse{
		Status:        -1,
		Message:       msg,
		DetailMessage: detail,
	})
}
