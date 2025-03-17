package main

import (
	"fmt"
	"gin-ranking/api/router"
)

func main() {
	// 获取路由器
	r := router.Router()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常：", err)
		}
	}()

	// 启动服务器
	r.Run(":9999")
}
