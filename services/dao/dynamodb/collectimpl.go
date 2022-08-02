package dynamodb

import "ramel-collectpage/models"

type CollectImplDynamoDB struct {
}

func (dao CollectImplDynamoDB) Create(u *models.Collect) error {
	return nil
}

func (dao CollectImplDynamoDB) GetAll() ([]models.Collect, error) {
	return nil, nil
}
