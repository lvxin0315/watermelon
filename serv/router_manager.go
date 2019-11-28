package serv

import (
	"errors"
	"github.com/Unknwon/goconfig"
	"github.com/lvxin0315/watermelon/helper"
	"github.com/lvxin0315/watermelon/local_db"
)

//路由处理监听
func RouterConfigRunner() {
	go initConfigRouter()
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

func FindGrpcApiUrlByRouter(router string) (string, error) {
	for _, config := range local_db.RouterConfigList {
		if config.Router == router {
			return config.GRPCApiUrl, nil
		}
	}
	return "", errors.New("undefined router")
}
