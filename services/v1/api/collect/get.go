package collect

import (
	"context"
	"encoding/json"
	"net/http"
	"ramel-collectpage/models"
	"ramel-collectpage/services/constants"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

func Get(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	userID := request.PathParameters[constants.F_NAME_USER_ID]

	filt := expression.Name(constants.F_NAME_USER_ID).Equal(expression.Value(userID))
	proj := expression.NamesList(
		expression.Name(constants.F_NAME_COLLECT_ID),
		expression.Name(constants.F_NAME_USER_ID),
		expression.Name(constants.F_NAME_URL),
	)

	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: http.StatusBadRequest,
		}, err
	}
	input := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(constants.T_NAME_COLLECT),
	}

	result, err := ddb.Scan(input)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: http.StatusBadRequest,
		}, err
	}

	items := []*models.Collect{}
	for _, i := range result.Items {
		item := models.Collect{}
		err = dynamodbattribute.UnmarshalMap(i, &item)
		if err != nil {
			return events.APIGatewayProxyResponse{
				Body:       err.Error(),
				StatusCode: http.StatusBadRequest,
			}, err
		}
		items = append(items, &item)
	}

	body, _ := json.Marshal(items)
	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: http.StatusOK,
	}, nil
}
