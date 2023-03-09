package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	username := "root"
	password := "Qwer1234"
	path := "192.168.1.153"
	port := "30830"
	dbname := "quiz"

	m := username + ":" + password + "@tcp(" + path + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       m,     // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据版本自动配置

	}

	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{}); err != nil {
		return nil
	} else {
		return db
	}
}
