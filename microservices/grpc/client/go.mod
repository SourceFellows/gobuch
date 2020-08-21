module golang.source-fellows.com/samples/grpc/client

go 1.13

require (
	golang.source-fellows.com/samples/grpc/hello v0.0.0
	google.golang.org/grpc v1.30.0
	google.golang.org/protobuf v1.25.0
)

replace golang.source-fellows.com/samples/grpc/hello => ./../hello
