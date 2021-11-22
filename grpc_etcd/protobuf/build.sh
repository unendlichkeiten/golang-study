#!/bin/bash
set -ue
cd /e/golangStudy/go-programming-tour-book/00-study/grpc_etcd
protoc -I. --go_out=plugins=grpc:. protobuf/pb/greet.proto