package behavioral

import "fmt"

type Strategyer interface {
	Execute(a, b int) int
}

type ConcreteStrategyAdd struct{}

func (s *ConcreteStrategyAdd) Execute(a, b int) int {
	return a + b
}

type ConcreteStrategySubtract struct{}

func (s *ConcreteStrategySubtract) Execute(a, b int) int {
	return a - b
}

type StrategyContext struct {
	strategy Strategyer
}

func (c *StrategyContext) SetStrategy(strategy Strategyer) {
	c.strategy = strategy
}

func (c *StrategyContext) ExecuteStrategy(a, b int) int {
	return c.strategy.Execute(a, b)
}

func Strategy() {
	context := &StrategyContext{}

	context.SetStrategy(&ConcreteStrategyAdd{})
	fmt.Println("10 + 5 =", context.ExecuteStrategy(10, 5))

	context.SetStrategy(&ConcreteStrategySubtract{})
	fmt.Println("10 - 5 =", context.ExecuteStrategy(10, 5))
}
