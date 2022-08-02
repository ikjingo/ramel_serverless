package utils

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func NewUUID(action string) string {
	return fmt.Sprintf("%s:%s", action, uuid.Must(uuid.NewV4(), nil).String())
}
