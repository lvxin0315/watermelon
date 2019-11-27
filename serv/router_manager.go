package serv

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type RouterConfig struct {
	Name       string
	Router     string
	Type       string
	GRPCApiUrl string
}

//路由处理监听
func RouterConfigRunner(engine *gin.Engine) {
	go func() {
		for true {
			time.Sleep(5 * time.Second)
			log.Println(engine.Routes())
		}
	}()
}

//读取配置静态结构体内容
