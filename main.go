package main

import (
	"fmt"
	"quiz/global"
	"quiz/initialize"
)

func main() {
	global.Global_DB = initialize.Gorm()
	if global.Global_DB == nil {
		fmt.Println("DB 失敗")
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
