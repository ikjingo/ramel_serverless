package dynamodb

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func get() *dynamodb.DynamoDB {
	region := os.Getenv("AWS_REGION")
	session, err := session.NewSession(&aws.Config{
		Region: &region,
	})
	if err != nil {
		log.Fatalln(fmt.Sprintf("Failed to connect to AWS: %s", err.Error()))
		return nil
	}
	return dynamodb.New(session) // Create DynamoDB client
}
