package main

import (
	"log"
	"net"

	"golang.source-fellows.com/samples/grpc/hello"
	"golang.source-fellows.com/samples/grpc/server"
	"google.golang.org/grpc"
)

func main() {
	srv := grpc.NewServer()
	hello.RegisterHelloWorldServiceServer(srv, &server.HelloService{})
	listener, _ := net.Listen("tcp", ":8080")
	log.Println("Starting Server ...")
	log.Fatal(srv.Serve(listener))
}
