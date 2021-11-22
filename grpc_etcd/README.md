### 项目目录
```shell script
./
├── cmd
│   └── grpc
│       ├── client_test.go # 测试文件
│       └── main.go # main文件
├── go.mod
├── go.sum
├── internal
│   ├── controller # grpc接口实现
│   │   └── foo_controller.go
│   ├── core
│   │   └── register # 服务注册实现
│   │       └── etcd.go
├── protobuf # proto原型文件和编译后的文件
│   ├── build.sh
│   ├── foo.proto
│   └── pb
│       └── foo.pb.go
└── README.md
```