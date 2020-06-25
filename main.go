package main

import (
	_ "github.com/gin-gonic/gin"
	"server/conf"
	"server/router"
)

func main() {
	// 读取配置
	conf.Init()

	// 装载路由
	r := router.NewRouter()
	_ = r.Run(":3000")
}
