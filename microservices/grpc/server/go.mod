module golang.source-fellows.com/samples/grpc/server

go 1.13

require (
	github.com/golang/protobuf v1.4.2
	golang.source-fellows.com/samples/grpc/hello v0.0.0
	google.golang.org/grpc v1.30.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v0.0.0-20200630190442-3de8449f8555 // indirect
	google.golang.org/protobuf v1.25.0
)

replace golang.source-fellows.com/samples/grpc/hello => ./../hello
