package main

import (
	"encoding/json"
	"fmt"
	"github.com/Kalimaha/hexa-go/pkg/repositories"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type SpamRequest struct {
	Spam string `json:"spam"`
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	data := SpamRequest{}

	var anyJson map[string]interface{}
	json.Unmarshal([]byte(request.Body), &anyJson)
	repositories.Ciaone(anyJson)

	json.Unmarshal([]byte(request.Body), &data)
	fmt.Printf("Original Request: %v\n", request.Body)
	fmt.Printf("Request: %v\n", data)
	fmt.Printf("Generic Request: %v\n", anyJson)

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       repositories.Ciao() + " from spam this time",
	}, nil
}

func main() {
	lambda.Start(handler)
}
