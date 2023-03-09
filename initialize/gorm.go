package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"quiz/model"
)

func Gorm() *gorm.DB {
	username := "root"
	password := "Qwer1234"
	path := "host.docker.internal"
	port := "30820"
	dbname := "quiz"

	dsn := username + ":" + password + "@tcp(" + path + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 字串預設長度
		SkipInitializeWithVersion: false, // 根據當前 MySQL 版本自動配置

	}

	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{}); err != nil {
		return nil
	} else {
		fmt.Println("********** MySQL Done! **********")
		return db
	}
}

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.Comment{},
	)
	if err != nil {
		fmt.Println("********** DB Error ! **********")
		os.Exit(0)
	}
}
