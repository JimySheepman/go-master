package common

import (
	"go-onion-architecture-sample/domain/model"
	"strconv"
)

func GetUserId(str string) (model.UserId, error) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return model.UserId(i), nil
}
