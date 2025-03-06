package common

import (
	"go-onion-architecture-sample/domain/model"
	"strconv"
)

func GetTodoId(str string) (model.TodoId, error) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return model.TodoId(i), nil
}
