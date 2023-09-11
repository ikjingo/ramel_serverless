package dynamodb

import (
	"ramel-collectpage/models"
	"ramel-collectpage/services/constants"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type CollectImplDynamoDB struct {
}

func (dao CollectImplDynamoDB) Create(collect *models.Collect) error {
	ddb := get()

	item, _ := dynamodbattribute.MarshalMap(collect)
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(constants.T_NAME_COLLECT),
	}
	_, err := ddb.PutItem(input)
	return err
}

func (dao CollectImplDynamoDB) GetAll(userID string) ([]*models.Collect, error) {
	ddb := get()

	filt := expression.Name(constants.F_NAME_USER_ID).Equal(expression.Value(userID))
	proj := expression.NamesList(
		expression.Name(constants.F_NAME_COLLECT_ID),
		expression.Name(constants.F_NAME_USER_ID),
		expression.Name(constants.F_NAME_URL),
	)
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	if err != nil {
		return nil, err
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
		return nil, err
	}

	items := []*models.Collect{}
	for _, i := range result.Items {
		item := models.Collect{}
		err = dynamodbattribute.UnmarshalMap(i, &item)
		if err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	return items, nil
}
