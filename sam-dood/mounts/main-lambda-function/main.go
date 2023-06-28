package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	"log"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (events.APIGatewayProxyResponse, error) {
	log.Println("Received : ", name.Name)
	return events.APIGatewayProxyResponse{
		Body:       "Hello " + name.Name,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
