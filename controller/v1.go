package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/lvxin0315/watermelon/serv"
	"log"
	"net/http"
)

func WsPage(c *gin.Context) {
	watermelonApi := c.Param("watermelon_api")
	//根据路由连接ws服务
	hub := serv.HubList[watermelonApi]
	upGrader := new(websocket.Upgrader)
	upGrader.CheckOrigin = func(r *http.Request) bool {
		//允许访问
		log.Println(r.RequestURI)
		return true
	}
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := serv.NewClient(hub, conn, make(chan []byte, 256))
	hub.Register <- client
	go client.ReadPump()
	go client.WritePump()
}

func HttpPage(c *gin.Context) {
	watermelonApi := c.Param("watermelon_api")
	c.JSON(200, gin.H{
		"message": watermelonApi,
	})
}
