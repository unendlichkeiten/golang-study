package main

import (
	"context"
	pb2 "go-programming-tour-book/golang-study/grpc_etcd/protobuf/pb"
	"testing"
	"time"

	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
)


func Test_Greet(t *testing.T) {

	cli, err := clientv3.NewFromURL("http://localhost:2379")
	if err != nil {
		t.Fatal(err)
	}
	builder, err := resolver.NewBuilder(cli)
	if err != nil {
		t.Fatal(err)
	}
	conn, err := grpc.Dial("etcd:///service/go-json_parser",
		grpc.WithResolvers(builder),
		grpc.WithBalancerName("round_robin"),
		grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	if err != nil {
		t.Fatal(err)
	}

	fooClient := pb2.NewFooClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	resp, err := fooClient.Greet(ctx, &pb2.GreetReq{
		MyName: "Bar",
		Msg:    "Hello, World",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp.Msg)
}