package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/router"
)

func InitGinRouter() *gin.Engine {
	// 创建gin服务
	ginServer := gin.Default()

	//提供服务组
	system := router.RouterGroupApp.System
	// 使用中间件
	ginServer.Use()

	// 业务模块接口
	privateGroup := ginServer.Group("/api")
	{
		system.InitUserRouter(privateGroup)
	}

	fmt.Println("router register success")
	return ginServer
}
