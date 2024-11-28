package main

import (
	"order-track/config"
	"order-track/routers"
	"order-track/utils"
)

func main() {
	// 初始化zap
	utils.InitLogger()
	// 初始化数据库
	config.InitDB()
	r := routes.SetupRouter()
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
