package main

import (
	"fmt"
	"github.com/Kalimaha/hexa-go/pkg/repositories"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Printf("<== == == == == == ==>\n")
	fmt.Printf("HTTP Method: %v\n", request.HTTPMethod)
	fmt.Printf("Headers: %v\n", request.Headers)
	fmt.Printf("Body: %v\n", request.Body)
	fmt.Printf("<== == == == == == ==>\n")

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       repositories.Ciao(),
	}, nil
}

func main() {
	lambda.Start(handler)
}
