package main

import (
	"fmt"
	"github.com/Kalimaha/hexa-go/pkg/repositories"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Print("Got request for '/.netlify/functions/hello', this message is dumpled by 'vendor/hello/pippo.go'")
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       repositories.Ciao(),
		//Body:       utils.IntroductionYourself("vendor/goodbye/pippo.go"),
	}, nil
}

func main() {
	lambda.Start(handler)
}
