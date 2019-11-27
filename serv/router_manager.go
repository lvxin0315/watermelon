package serv

import (
	"github.com/Unknwon/goconfig"
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/watermelon/helper"
	"github.com/lvxin0315/watermelon/local_db"
	"log"
	"time"
)

//路由处理监听
func RouterConfigRunner(engine *gin.Engine) {
	initConfigRouter()
	go func() {
		for true {
			time.Sleep(5 * time.Second)
			log.Println(engine.Routes())
		}
	}()
}

//使用配置文件生成静态结构体内容
func initConfigRouter() {
	cfg, err := goconfig.LoadConfigFile("etc/router.ini")
	helper.CheckError("InitConfigRouter", err)
	for _, sectionName := range cfg.GetSectionList() {
		configMap, err := cfg.GetSection(sectionName)
		helper.CheckError("cfg.GetSection", err)
		local_db.RouterConfigList = append(local_db.RouterConfigList, &local_db.RouterConfig{
			Name:       sectionName,
			Router:     configMap["router"],
			Type:       configMap["type"],
			GRPCApiUrl: configMap["grpc_api_url"],
		})
		//websocket开启服务
		if configMap["type"] == "ws" {
			NewHub(configMap["router"]).Run()
		}
	}
}

//读取配置静态结构体内容
func UpdateStaticStruct() {
	//比较是否有变化
	//reflect.DeepEqual(a, b)
}
