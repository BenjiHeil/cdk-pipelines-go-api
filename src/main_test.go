package main

import (
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandlerWithGoodParams(t *testing.T) {

	//arrange
	x, y, want := 16, 2, 8
	request := events.APIGatewayProxyRequest{
		Path: "/divide",
		QueryStringParameters: map[string]string{
			"x": fmt.Sprint(x),
			"y": fmt.Sprint(y),
		},
	}

	//act
	response, err := handler(nil, request)

	//assert
	if err != nil {
		t.Fatalf("handler returned err %v", err)
	}

	if response.Body != fmt.Sprint(want) {
		t.Fatalf("handler returned incorrect value")
	}
}

func TestHandlerWithDivideByZero(t *testing.T) {

	//arrange
	x, y := 16, 0
	request := events.APIGatewayProxyRequest{
		Path: "/divide",
		QueryStringParameters: map[string]string{
			"x": fmt.Sprint(x),
			"y": fmt.Sprint(y),
		},
	}

	//act
	response, err := handler(nil, request)

	//assert
	if err != nil {
		t.Fatalf("handler returned err %v", err)
	}

	if response.StatusCode < 400 {
		t.Fatalf("handler did not return a bad request for divide by zero")
	}
}
