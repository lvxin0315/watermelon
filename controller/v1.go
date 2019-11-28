package controller

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	wmgrpc "github.com/lvxin0315/watermelon/grpc"
	"github.com/lvxin0315/watermelon/serv"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
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
		logrus.Errorln(err)
		return
	}
	client := serv.NewClient(hub, conn, make(chan []byte, 256))
	hub.Register <- client
	go client.ReadPump()
	go client.WritePump()
}

func HttpPage(c *gin.Context) {
	watermelonApi := c.Param("watermelon_api")
	address, err := serv.FindGrpcApiUrlByRouter(watermelonApi)
	if err != nil {
		logrus.Errorln("HttpPage.serv.FindGrpcApiUrlByRouter:", err)
		c.JSON(200, err.Error())
		return
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		logrus.Errorln("grpc.Dial:", err)
		c.JSON(200, err.Error())
		return
	}
	client := wmgrpc.NewHttpServerClient(conn)
	result, err := client.GetData(context.Background(), &wmgrpc.Empty{})
	if err != nil {
		logrus.Errorln("client.GetData:", err)
		c.JSON(200, err.Error())
		return
	}
	data := make(map[string]interface{})
	err = json.Unmarshal(result.Data, &data)
	if err != nil {
		logrus.Errorln("json.Unmarshal:", err)
		c.JSON(200, err.Error())
		return
	}
	c.JSON(200, data)
	return
}
