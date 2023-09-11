package collect

import (
	"context"
	"encoding/json"
	"net/http"
	"ramel-collectpage/models"
	"ramel-collectpage/services/constants"
	"ramel-collectpage/services/dao/factory"
	"ramel-collectpage/services/utils"

	"github.com/aws/aws-lambda-go/events"
)

func Post(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	collect := &models.Collect{}
	if err := json.Unmarshal([]byte(request.Body), collect); err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: http.StatusBadRequest,
		}, err
	}
	collect.CollectID = utils.NewUUID(constants.Collect)
	collect.UserID = request.PathParameters[constants.F_NAME_USER_ID]

	collectDao := factory.FactoryDao("dynamodb") // 설정값에 따라 DB 설정 하도록 수정 필요

	if err := collectDao.Create(collect); err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       collect.CollectID,
		StatusCode: http.StatusOK,
	}, nil
}
