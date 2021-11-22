package controller

import (
	"context"
	"fmt"
	pb2 "go-programming-tour-book/golang-study/grpc_etcd/protobuf/pb"
)

type FooController struct {
}

func NewFooController() *FooController {
	f := &FooController{}
	return f
}

func (f *FooController) Greet(ctx context.Context, in *pb2.GreetReq) (*pb2.GreetResp, error) {
	reply := fmt.Sprintf("Hello %s, I got your msg:%s", in.GetMyName(), in.GetMsg())
	out := &pb2.GreetResp{}
	out.Msg = reply
	return out, nil
}