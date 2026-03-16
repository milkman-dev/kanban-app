package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func handler(ctx context.Context, request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyStreamingResponse, error) {
	h := HealthResponse{Status: "OK"}

	body, err := json.Marshal(h)
	if err != nil {
		return events.APIGatewayProxyStreamingResponse{}, err
	}

	return events.APIGatewayProxyStreamingResponse{
		StatusCode: http.StatusOK,
		Body:       bytes.NewReader(body),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil

}

func main() {

}
