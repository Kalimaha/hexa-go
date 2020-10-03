package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"spam"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Print("Got request for '/.netlify/functions/hello', this message is dumpled by 'cmd/hello/pippo.go'")
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       spam.SuggestShit(),
		//Body:       utils.IntroductionYourself("cmd/goodbye/pippo.go"),
	}, nil
}

func main() {
	lambda.Start(handler)
}
