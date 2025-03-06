package creational

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

func AnimalFactory(animalType string) Animal {
	if animalType == "dog" {
		return Dog{}
	}
	if animalType == "cat" {
		return Cat{}
	}
	return nil
}

func FactoryMethod() {
	animal1 := AnimalFactory("dog")
	fmt.Println(animal1.Speak())

	animal2 := AnimalFactory("cat")
	fmt.Println(animal2.Speak())
}
