package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/watermelon/controller"
	"github.com/lvxin0315/watermelon/serv"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//TODO
	hub := serv.NewHub("/ws1")

	go func() {
		hub.Run()
	}()

	go func() {
		time.Sleep(3 * time.Second)
		hub.Broadcast([]byte("a"))
	}()

	r.GET("/ws1", controller.WsPage)

	r.Run()
}
