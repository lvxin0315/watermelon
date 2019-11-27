package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/watermelon/controller"
	"github.com/lvxin0315/watermelon/serv"
	"time"
)

func main() {
	engine := gin.Default()
	//配合检查用
	engine.GET("/ping", controller.Ping)

	//TODO
	hub := serv.NewHub("/ws1")

	go func() {
		hub.Run()
	}()

	go func() {
		time.Sleep(3 * time.Second)
		hub.Broadcast([]byte("a"))
	}()

	engine.GET("/ws1", controller.WsPage)

	//路由处理
	serv.RouterConfigRunner(engine)

	engine.Run()
}
