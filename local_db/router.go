package local_db

type RouterConfig struct {
	Name       string
	Router     string
	Type       string
	GRPCApiUrl string
}

var RouterConfigList []*RouterConfig
