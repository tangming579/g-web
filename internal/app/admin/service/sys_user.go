package service

import (
	"go-web/internal/db"
	syslogger "go-web/internal/logger"
	"go-web/internal/models"
)

func GetUser(userId string) (model *models.SysUser) {
	syslogger.Info("get")
	db.Datasource["postgres"].DB.Table("sys_user").Where("\"usercode\"=?", userId).Take(&model)
	return
}

func DeleteUser(userId string) (model *models.SysUser) {
	syslogger.Info("delete")
	db.Datasource["postgres"].DB.Table("sys_user").Where("\"usercode\"=?", userId).Delete(&model)
	return
}

func CreateUser(userInfo models.SysUser) {
	syslogger.Info("create")
	db.Datasource["postgres"].DB.Table("sys_user").Create(&userInfo)
}

func UpdateUser(userInfo models.SysUser) {
	syslogger.Info("update")
	db.Datasource["postgres"].DB.Table("sys_user").Where("\"usercode\"=?", userInfo.Usercode).
		Save(&userInfo)
}
