package util

import (
	uuid "github.com/satori/go.uuid"
)

func getUuid() string {

	v4, _ := uuid.NewV4()

	return v4.String()
}
