package service

import (
	"go-web/internal/db"
	syslogger "go-web/internal/logger"
	"go-web/internal/models"
)

func GetUser(userId string) (model *models.SysUser) {
	syslogger.Info("get")
	db.DB.Table("sys_user").Where("\"usercode\"=?", userId).Take(&model)
	return
}

func GetUserList(pageSize int, pageNum int) (users *[]models.SysUser, total int64) {
	syslogger.Info("getList")
	db.DB.Table("sys_user").Count(&total)
	db.DB.Table("sys_user").Limit(pageSize).Offset((pageNum - 1) * pageSize).
		Find(&users)
	return
}

func DeleteUser(userId string) (model *models.SysUser) {
	syslogger.Info("delete")
	db.DB.Table("sys_user").Where("\"usercode\"=?", userId).Delete(&model)
	return
}

func CreateUser(userInfo models.SysUser) {
	syslogger.Info("create")
	db.DB.Table("sys_user").Create(&userInfo)
}

func UpdateUser(userInfo models.SysUser) {
	syslogger.Info("update")
	db.DB.Table("sys_user").Where("\"usercode\"=?", userInfo.Usercode).
		Save(&userInfo)
}
