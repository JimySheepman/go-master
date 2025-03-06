package structural

import "fmt"

type Coffee interface {
	GetCost() int
	GetDescription() string
}

type SimpleCoffee struct{}

func (c *SimpleCoffee) GetCost() int {
	return 5
}

func (c *SimpleCoffee) GetDescription() string {
	return "Simple coffee"
}

type MilkDecorator struct {
	coffee Coffee
}

func (m *MilkDecorator) GetCost() int {
	return m.coffee.GetCost() + 2
}

func (m *MilkDecorator) GetDescription() string {
	return m.coffee.GetDescription() + ", milk"
}

type SugarDecorator struct {
	coffee Coffee
}

func (s *SugarDecorator) GetCost() int {
	return s.coffee.GetCost() + 1
}

func (s *SugarDecorator) GetDescription() string {
	return s.coffee.GetDescription() + ", sugar"
}

func Decorator() {
	coffee := &SimpleCoffee{}
	fmt.Println(coffee.GetDescription(), "$", coffee.GetCost())

	coffeeWithMilk := &MilkDecorator{coffee: coffee}
	fmt.Println(coffeeWithMilk.GetDescription(), "$", coffeeWithMilk.GetCost())

	coffeeWithMilkAndSugar := &SugarDecorator{coffee: coffeeWithMilk}
	fmt.Println(coffeeWithMilkAndSugar.GetDescription(), "$", coffeeWithMilkAndSugar.GetCost())
}
