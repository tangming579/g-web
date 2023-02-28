package service

import (
	"go-web/internal/db"
	"go-web/internal/models"
)

func Get(userId string) (model *models.SysUser) {
	db.Datasource["postgres"].DB.Table("sys_user").Where("\"usercode\"=?", userId).Take(&model)
	return
}
