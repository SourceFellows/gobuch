package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {

	lc, _ := lambdacontext.FromContext(ctx)

	return fmt.Sprintf("Hello %s! %s", name.Name, lc.AwsRequestID), nil
}

func main() {
	lambda.Start(HandleRequest)
}
