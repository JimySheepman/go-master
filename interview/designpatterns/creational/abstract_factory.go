package creational

import "fmt"

type Shoe interface {
	Size() int
}

type Shirt interface {
	Color() string
}

type AdidasShoe struct{}

func (a AdidasShoe) Size() int {
	return 42
}

type AdidasShirt struct{}

func (a AdidasShirt) Color() string {
	return "Red"
}

type NikeShoe struct{}

func (n NikeShoe) Size() int {
	return 44
}

type NikeShirt struct{}

func (n NikeShirt) Color() string {
	return "Blue"
}

type SportsFactory interface {
	MakeShoe() Shoe
	MakeShirt() Shirt
}

type AdidasFactory struct{}

func (a AdidasFactory) MakeShoe() Shoe {
	return AdidasShoe{}
}

func (a AdidasFactory) MakeShirt() Shirt {
	return AdidasShirt{}
}

type NikeFactory struct{}

func (n NikeFactory) MakeShoe() Shoe {
	return NikeShoe{}
}

func (n NikeFactory) MakeShirt() Shirt {
	return NikeShirt{}
}

func GetSportsFactory(brand string) SportsFactory {
	if brand == "adidas" {
		return AdidasFactory{}
	}
	if brand == "nike" {
		return NikeFactory{}
	}
	return nil
}

func AbstractFactory() {
	adidasFactory := GetSportsFactory("adidas")
	nikeFactory := GetSportsFactory("nike")

	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()
	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()

	fmt.Println("Adidas Shoe Size:", adidasShoe.Size())
	fmt.Println("Adidas Shirt Color:", adidasShirt.Color())
	fmt.Println("Nike Shoe Size:", nikeShoe.Size())
	fmt.Println("Nike Shirt Color:", nikeShirt.Color())
}
