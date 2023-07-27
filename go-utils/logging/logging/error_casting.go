package main

import (
	"fmt"

	"github.com/JimySheepman/go-master/go-utils/logging/api"
)

func errorCasting() {
	err := api.NewCustomError(api.ErrInternalServer, "", map[string]interface{}{
		"Description": "description error",
		"ResultType":  "result type error",
		"testVAlue":   "test value",
	})
	_ = err

	err2 := fmt.Errorf("test error")

	a, ok := err2.(api.Error)
	if !ok {
		fmt.Println("cast error 1. ")
		return
	}

	fmt.Println("Description: ", a.Params["Description"])
	fmt.Println("ResultType: ", a.Params["ResultType"])

	e, ok := a.Params["test"]
	if !ok {
		fmt.Println("cast error 2. ")
		return
	}

	fmt.Println("test done! : ", e)
}
