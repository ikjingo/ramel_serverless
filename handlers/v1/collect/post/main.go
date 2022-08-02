package main

import (
	"ramel-collectpage/services/v1/api/collect"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(collect.Post)
}
