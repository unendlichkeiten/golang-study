clean:
	rm pb/*

gen:
	protoc --plugin=protoc-gen-go=/d/workspace/golang/bin/protoc-gen-go.exe --go_out=plugins=grpc:pb --proto_path=proto proto/*.proto