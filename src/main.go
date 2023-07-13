package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var response events.APIGatewayProxyResponse
	switch event.Path {
	case "/divide":
		{
			x, _ := strconv.Atoi(event.QueryStringParameters["x"])
			y, _ := strconv.Atoi(event.QueryStringParameters["y"])
			dividend, err := divide(x, y)
			if err != nil {
				response = events.APIGatewayProxyResponse{
					StatusCode: 400,
					Body:       err.Error(),
				}
			} else {
				response = events.APIGatewayProxyResponse{
					StatusCode: 200,
					Body:       fmt.Sprint(dividend),
				}
			}
		}
	default:
		{
			response = events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       "\"Hello from Lambda!\"",
			}
		}
	}
	return response, nil
}

func main() {
	lambda.Start(handler)
}

type DivideByZeroError struct {
	params map[string]int
}

func (d *DivideByZeroError) Error() string {
	return fmt.Sprintf("Cannot divide by zero: Operands %v", d.params)
}

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, &DivideByZeroError{params: map[string]int{"x": x, "y": y}}
	}
	return x / y, nil
}
