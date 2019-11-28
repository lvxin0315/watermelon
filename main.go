package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/watermelon/controller"
	"github.com/lvxin0315/watermelon/serv"
)

func main() {
	engine := gin.Default()
	//配合检查用
	engine.GET("/ping", controller.Ping)
	//http
	engine.Any("/v1/h/:watermelon_api", controller.HttpPage)
	//ws
	engine.Any("/v1/w/:watermelon_api", controller.WsPage)

	//路由处理
	serv.RouterConfigRunner()

	engine.Run()
}
