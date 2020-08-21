package server

import (
	context "context"

	empty "github.com/golang/protobuf/ptypes/empty"
	"golang.source-fellows.com/samples/grpc/hello"
)

type HelloService struct {
	hello.UnimplementedHelloWorldServiceServer
}

func (hs *HelloService) SayHello(context.Context, *empty.Empty) (*hello.HelloMessage, error) {
	msg := &hello.HelloMessage{MessageText: "Hello World"}
	return msg, nil
}
