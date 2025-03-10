package main

import (
	"go-generics/02/basic"
	"go-generics/02/behaviorconstraint"
	"go-generics/02/channels"
	"go-generics/02/hashtables"
	"go-generics/02/multitypeparameters"
	"go-generics/02/sliceconstraints"
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
	multitypeparameters.Example,
	sliceconstraints.Example,
	channels.Example,
	hashtables.Example,
}

func main() {
	for _, run := range examples {
		run()
	}
}
