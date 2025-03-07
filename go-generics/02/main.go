package main

import (
	"go-generics/02/basic"
	"go-generics/02/behaviorconstraint"
	"go-generics/02/structtypes"
	"go-generics/02/typeconstraint"
	"go-generics/02/underlyingtypes"
)

type example func()

var examples = []example{
	basic.Example,
	underlyingtypes.Example,
	structtypes.Example,
	behaviorconstraint.Example,
	typeconstraint.Example,
}

func main() {
	for _, ex := range examples {
		ex()
	}
}
