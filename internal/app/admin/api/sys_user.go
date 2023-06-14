package api

import (
	"github.com/gin-gonic/gin"
	"go-web/internal/app/admin/service"
	"go-web/internal/http/response"
	sysLogger "go-web/internal/logger"
	"go-web/internal/models"
	"strconv"
)

func GetUserInfo(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		response.Failure(c, "userId不能为空")
		return
	}
	sysLogger.Info("get user")
	//object := service.GetUser(userId)
	response.Ok(c, nil)
	return
}

func GetUserList(c *gin.Context) {
	pageNum, err := strconv.Atoi(c.Query("pageNum"))
	if err != nil {
		pageNum = 1
	}
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil {
		pageSize = 10
	}

	users, total := service.GetUserList(pageSize, pageNum)
	response.Ok(c, gin.H{"list": users, "total": total})
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
