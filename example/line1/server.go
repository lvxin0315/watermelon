package main

import (
	"context"
	"encoding/json"
	"github.com/lvxin0315/watermelon/data_model"
	"github.com/lvxin0315/watermelon/grpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	PORT = ":50001"
)

type server struct{}

func (s server) GetData(context.Context, *wmgrpc.Empty) (*wmgrpc.Result, error) {
	lineModel := new(data_model.LineData)
	grpcJson := new(wmgrpc.Result)
	jsonBytes, err := json.Marshal(&lineModel)
	if err != nil {
		return grpcJson, err
	}
	grpcJson.Data = jsonBytes
	return grpcJson, nil
}

func main() {
	lis, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	wmgrpc.RegisterHttpServerServer(s, &server{})
	log.Println("rpc服务已经开启,port", PORT)
	err = s.Serve(lis)
	log.Fatal(err)
}
