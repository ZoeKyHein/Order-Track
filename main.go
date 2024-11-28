package main

import (
	"order-track/config"
	"order-track/routers"
)

func main() {

	// 初始化数据库
	config.InitDB()
	r := routes.SetupRouter()
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
