package main

import (
	"context"
	"fmt"
	controller2 "go-programming-tour-book/golang-study/grpc_etcd/internal/controller"
	register2 "go-programming-tour-book/golang-study/grpc_etcd/internal/core/register"
	pb2 "go-programming-tour-book/golang-study/grpc_etcd/protobuf/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

func registerService(s *grpc.Server) {
	pb2.RegisterFooServer(s, controller2.NewFooController())
}

func main() {
	addr := "127.0.0.1:8080"
	ctx := context.Background()

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	registerService(s)

	if err := register2.Register(ctx, "go-json_parser", addr); err != nil { // 服务注册名: go-json_parser
		log.Fatalf("register %s failed:%v", "go-json_parser", err)
	}

	fmt.Printf("start grpc server:%s", addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}