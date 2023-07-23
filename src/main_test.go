package main

import (
	"fmt"
	"math/rand"
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
	response, _ := handler(nil, request)

	if response.StatusCode < 400 {
		t.Fatalf("handler did not return a bad request for divide by zero")
	}
}

func TestHandlerWithNoX(t *testing.T) {
	//arrange
	y := 16
	request := events.APIGatewayProxyRequest{
		Path: "/divide",
		QueryStringParameters: map[string]string{
			"y": fmt.Sprint(y),
		},
	}

	//act
	response, _ := handler(nil, request)

	if response.StatusCode < 400 {
		t.Fatalf("handler did not return a bad request for missing x")
	}
}

func TestHandlerWithNoY(t *testing.T) {
	//arrange
	x := 16
	request := events.APIGatewayProxyRequest{
		Path: "/divide",
		QueryStringParameters: map[string]string{
			"x": fmt.Sprint(x),
		},
	}

	//act
	response, _ := handler(nil, request)

	if response.StatusCode < 400 {
		t.Fatalf("handler did not return a bad request for missing y")
	}
}

func BenchmarkDivide(b *testing.B) {
	defer func() { _ = recover() }()
	for i := 0; i < b.N; i++ {
		//arrange
		x, y := rand.Int(), rand.Int()

		//act
		_, _ = Divide(x, y)
	}
}

func BenchmarkHandler(b *testing.B) {
	defer func() { _ = recover() }()
	for i := 0; i < b.N; i++ {
		//arrange
		x, y := rand.Int(), rand.Int()

		request := events.APIGatewayProxyRequest{
			Path: "/divide",
			QueryStringParameters: map[string]string{
				"x": fmt.Sprint(x),
				"y": fmt.Sprint(y),
			},
		}

		//act
		_, _ = handler(nil, request)
	}
}
