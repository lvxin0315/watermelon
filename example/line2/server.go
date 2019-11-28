package main

import (
	"encoding/json"
	"github.com/lvxin0315/watermelon/data_model"
	"github.com/lvxin0315/watermelon/grpc"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const (
	PORT = ":50003"
)

type server struct{}

func (s server) GetData(stream wmgrpc.WebSocketServer_GetDataServer) error {
	for true {
		time.Sleep(time.Second)
		lineModel := setData()
		jsonBytes, _ := json.Marshal(lineModel)
		log.Println(string(jsonBytes))
		if err := stream.Send(&wmgrpc.Result{Data: jsonBytes}); err != nil {
			log.Println(err)
		}
	}
	return nil
}

func setData() *data_model.LineData {
	lineModel := new(data_model.LineData)
	lineModel.Grid = &data_model.Grid{
		Bottom:       "3%",
		ContainLabel: false,
		Left:         "3%",
		Right:        "4&",
	}
	lineModel.Legend = &data_model.Legend{Data: []string{"Y轴1", "Y轴2", "Y轴3"}}

	lineModel.Series = append(lineModel.Series, &data_model.Series{
		Data:  []int64{120, 132, 101, 134, 90, 230, 210},
		Name:  "Y轴1",
		Stack: "总量",
		Type:  "line",
	})

	lineModel.Series = append(lineModel.Series, &data_model.Series{
		Data:  []int64{220, 182, 191, 234, 290, 330, 310},
		Name:  "Y轴2",
		Stack: "总量",
		Type:  "line",
	})

	lineModel.Series = append(lineModel.Series, &data_model.Series{
		Data:  []int64{150, 232, 201, 154, 190, 330, 410},
		Name:  "Y轴3",
		Stack: "总量",
		Type:  "line",
	})

	lineModel.Title = &data_model.Title{Text: "测试折线图"}

	lineModel.XAxis = &data_model.XAxis{
		BoundaryGap: false,
		Data:        []string{"X轴1", "X轴2", "X轴3", "X轴4", "X轴5", "X轴6", "X轴7"},
		Type:        "category",
	}

	lineModel.YAxis = &data_model.YAxis{Type: "value"}
	return lineModel
}

func main() {
	lis, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	wmgrpc.RegisterWebSocketServerServer(s, &server{})
	log.Println("rpc服务已经开启,port", PORT)
	err = s.Serve(lis)
	log.Fatal(err)
}
