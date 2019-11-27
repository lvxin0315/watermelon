package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/lvxin0315/watermelon/serv"
	"log"
	"net/http"
)

func WsPage(c *gin.Context) {
	hub := serv.HubList[c.Request.RequestURI]
	upGrader := new(websocket.Upgrader)
	upGrader.CheckOrigin = func(r *http.Request) bool {
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
