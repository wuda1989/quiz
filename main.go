package main

import (
	"database/sql"
	"fmt"
	"quiz/global"
	"quiz/initialize"
)

func main() {
	global.Global_DB = initialize.Gorm()
	if global.Global_DB != nil {
		initialize.RegisterTables(global.Global_DB) // 初始化資料表
		db, _ := global.Global_DB.DB()
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				fmt.Println("DB尚未初始化")
				return
			}
		}(db)
	} else {
		fmt.Println("DB尚未初始化")
		return
	}

	RunServer()
}

func RunServer() {
	Router := initialize.Routers()
	err := Router.Run(":8888")
	if err != nil {
		fmt.Println("啟動失敗")
		return
	}
	fmt.Println("監聽:8888")
}
