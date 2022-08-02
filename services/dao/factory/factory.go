package factory

import (
	"log"
	"ramel-collectpage/services/dao/dynamodb"
	"ramel-collectpage/services/dao/interfaces"
)

func FactoryDao(e string) interfaces.UserDao {
	var i interfaces.UserDao
	switch e {
	case "dynamodb":
		i = dynamodb.CollectImplDynamoDB{}
	default:
		log.Fatalf("El motor %s no esta implementado", e)
		return nil
	}

	return i
}
