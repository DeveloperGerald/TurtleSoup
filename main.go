package main

import (
	"fmt"

	"github.com/DeveloperGerald/TurtleSoup/config"
	"github.com/DeveloperGerald/TurtleSoup/repository"
	"github.com/DeveloperGerald/TurtleSoup/router"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	repository.Init()

	// 创建一个默认的 Gin 路由引擎
	r := gin.Default()

	// 注册路由
	router.RegisterRouter(r)

	// 启动服务
	r.Run(fmt.Sprintf(":%s", config.GetConfig().Server.Port)) // 默认在 8080 端口启动
}
