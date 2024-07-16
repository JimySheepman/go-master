package behavioral

import "fmt"

type Stater interface {
	Handle(context *StateContext)
}

type StateContext struct {
	state Stater
}

func (c *StateContext) SetState(state Stater) {
	c.state = state
}

func (c *StateContext) Request() {
	c.state.Handle(c)
}

type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle(context *StateContext) {
	fmt.Println("State A handling request and transitioning to State B")
	context.SetState(&ConcreteStateB{})
}

type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle(context *StateContext) {
	fmt.Println("State B handling request and transitioning to State A")
	context.SetState(&ConcreteStateA{})
}

func State() {
	context := &StateContext{state: &ConcreteStateA{}}

	context.Request()
	context.Request()
	context.Request()
	context.Request()
}
