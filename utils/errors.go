package utils

import (
	"errors"
	"fmt"
)

func FailedMessage(message string, err error) (string, error) {
	fmt.Println(err)

	return "", errors.New(message)
}
