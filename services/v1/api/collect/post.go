package collect

import (
	"context"
	"encoding/json"
	"ramel-collectpage/models"
	"ramel-collectpage/services/constants"
	"ramel-collectpage/services/utils"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func Post(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	collect := &models.Collect{}
	if err := json.Unmarshal([]byte(request.Body), collect); err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 400,
		}, err
	}

	collect.CollectID = utils.NewUUID(constants.Collect)
	collect.UserID = request.PathParameters["user_id"]
	item, _ := dynamodbattribute.MarshalMap(collect)
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("Collect"),
	}
	if _, err := ddb.PutItem(input); err != nil {
		return events.APIGatewayProxyResponse{ // Error HTTP response
			Body:       err.Error(),
			StatusCode: 500,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       collect.CollectID,
		StatusCode: 200,
	}, nil
}
