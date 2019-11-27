module github.com/lvxin0315/watermelon

go 1.12

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190927123631-a832865fa7ad
	golang.org/x/net => github.com/golang/net v0.0.0-20190926025831-c00fd9afed17
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190911185100-cd5d95a43a6e
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190927073244-c990c680b611
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190929041059-e7abfedfabcf
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20190717185122-a985d3407aa7
)

require (
	github.com/Unknwon/goconfig v0.0.0-20191126170842-860a72fb44fd
	github.com/gin-gonic/gin v1.5.0
	github.com/gorilla/websocket v1.4.1
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/sirupsen/logrus v1.4.2
)
