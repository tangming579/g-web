package api

import (
	"github.com/gin-gonic/gin"
	"go-web/internal/app/admin/service"
	"go-web/internal/http/response"
	sysLogger "go-web/internal/logger"
	"go-web/internal/models"
)

func GetUserInfo(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		response.Failure(c, "userId不能为空")
		return
	}
	object := service.GetUser(userId)
	response.Ok(c, object)
	return
}

func UpdateUser(c *gin.Context) {
	var reqInfo models.SysUser
	err := c.BindJSON(&reqInfo)
	if err != nil {
		sysLogger.Info(err.Error())
		response.Failure(c, "post data error")
		return
	} else {
		service.UpdateUser(reqInfo)
		response.Ok(c, nil)
		return
	}
}

func DeleteUser(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		response.Failure(c, "userId不能为空")
		return
	}
	service.DeleteUser(userId)
	response.Ok(c, nil)
	return
}

func CreateUser(c *gin.Context) {
	var reqInfo models.SysUser
	err := c.BindJSON(&reqInfo)
	if err != nil {
		sysLogger.Info(err.Error())
		response.Failure(c, "post data error")
		return
	} else {
		service.CreateUser(reqInfo)
		response.Ok(c, nil)
		return
	}
}
