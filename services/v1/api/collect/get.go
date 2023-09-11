package collect

import (
	"context"
	"encoding/json"
	"net/http"
	"ramel-collectpage/services/constants"
	"ramel-collectpage/services/dao/factory"

	"github.com/aws/aws-lambda-go/events"
)

func Get(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	userID := request.PathParameters[constants.F_NAME_USER_ID]

	collectDao := factory.FactoryDao("dynamodb") // 설정값에 따라 DB 설정 하도록 수정 필요
	items, err := collectDao.GetAll(userID)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: http.StatusBadRequest,
		}, err
	}

	body, _ := json.Marshal(items)
	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: http.StatusOK,
	}, nil
}
