%go mod init github.com/thomasmatt88/go_protocol_buffer_test
Add import "github.com/golang/protobuf/proto" to main.go
%go mod tidy

given schema definition(ie search_request.proto) you can now generate the classes you'll need to
read and write SearchRequest messages using the protocol buffer compiler protoc on search_request.proto

Install protoc
%brew install protobuf
%protoc --version
libprotoc 3.21.12

The protocol buffer compiler requires a plugin to generate Go code.
It augments the protoc compiler so that it knows how to generate Go specific code for a given .proto file
install protocol compiler plugin for Go
This will install a protoc-gen-go binary in $GOBIN. Set the $GOBIN environment variable to change the installation location
% go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
% ls $GOPATH/bin
protoc-gen-go
% $GOPATH/bin/protoc-gen-go --version
protoc-gen-go v1.28.1

$GOPATH/bin or $GOBIN must be in your $PATH for the protocol buffer compiler to find it

at root of this repo
%protoc --go_out="./go_out" ./model/search_request.proto
This will autogenerate go package and code in go_out directory

$ go run main.go
