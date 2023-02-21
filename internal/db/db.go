package db

import (
	"fmt"
	"go-web/internal/config"

	"go-web/internal/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseMap map[string]*Database

type Database struct {
	name     string
	dbType   string
	host     string
	port     int
	dbname   string
	username string
	password string
	DB       *gorm.DB
}

var Datasource DatabaseMap

func init() {
	Datasource = make(DatabaseMap)
	dbconfig := config.AppConfig.Get("app.database")
	for _, dbinfo := range dbconfig.([]interface{}) {
		name := dbinfo.(map[string]interface{})["name"].(string)
		dbType := dbinfo.(map[string]interface{})["type"].(string)
		host := dbinfo.(map[string]interface{})["host"].(string)
		port := dbinfo.(map[string]interface{})["port"].(int)
		dbname := dbinfo.(map[string]interface{})["dbname"].(string)
		username := dbinfo.(map[string]interface{})["username"].(string)
		password := dbinfo.(map[string]interface{})["password"].(string)
		db := Database{
			name:     name,
			dbType:   dbType,
			host:     host,
			port:     port,
			dbname:   dbname,
			username: username,
			password: password,
		}
		Datasource[name] = &db
	}
}
func InitDB() {
	logger.Info("(启动中) 初始化数据库...")
	var err error
	for dbname, dbinfo := range Datasource {
		switch dbinfo.dbType {
		case "postgres":
			err = initPg(dbinfo)
		}
		if err != nil {
			panic(fmt.Errorf("初始化数据库:%s 失败: %v", dbname, err))
		}
	}
}
func initPg(dbinfo *Database) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		dbinfo.host,
		dbinfo.username,
		dbinfo.password,
		dbinfo.dbname,
		dbinfo.port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	dbinfo.DB = db
	return err
}
