package main

import (
	"fmt"
	"quiz/global"
	"quiz/initialize"
)

func main() {
	//time.Sleep(20 * time.Second) // 等待mysql container 完成

	global.Global_DB = initialize.Gorm()
	if global.Global_DB != nil {
		initialize.RegisterTables(global.Global_DB) // 初始化資料表
		db, _ := global.Global_DB.DB()
		defer db.Close()
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
