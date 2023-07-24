package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch event.Path {
	case "/divide":
		{
			var x, y, dividend int
			var err error
			if x, err = strconv.Atoi(event.QueryStringParameters["x"]); err != nil {
				return events.APIGatewayProxyResponse{
					StatusCode: 400,
					Body:       MissingParamError{params: event.QueryStringParameters}.Error(),
				}, nil
			}

			if y, err = strconv.Atoi(event.QueryStringParameters["y"]); err != nil {
				return events.APIGatewayProxyResponse{
					StatusCode: 400,
					Body:       MissingParamError{params: event.QueryStringParameters}.Error(),
				}, nil
			}

			if dividend, err = Divide(x, y); err != nil {
				return events.APIGatewayProxyResponse{
					StatusCode: 400,
					Body:       err.Error(),
				}, nil
			}

			return events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       fmt.Sprint(dividend),
			}, nil

		}
	default:
		{
			return events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       "\"Hello from Lambda!\"",
			}, nil
		}
	}
}

func main() {
	lambda.Start(handler)
}

type DivideByZeroError struct {
	params map[string]int
}

type MissingParamError struct {
	params map[string]string
}

func (m MissingParamError) Error() string {
	return fmt.Sprintf("Missing required parameters x or y. params: %v", m.params)
}

func (d *DivideByZeroError) Error() string {
	return fmt.Sprintf("Cannot divide by zero: Operands %v", d.params)
}

func Divide(x, y int) (int, error) {
	time.Sleep(2000)
	if y == 0 {
		return 0, &DivideByZeroError{params: map[string]int{"x": x, "y": y}}
	}
	return x / y, nil
}
