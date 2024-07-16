package oop

import "fmt"

type oopFunc func()

var oops = []oopFunc{
	Encapsulation,
	Inheritance,
	Polymorphism,
	Abstraction,
	ConstructorDestructor,
}

var oopsName = map[int]string{
	0: "Encapsulation",
	1: "Inheritance",
	2: "Polymorphism",
	3: "Abstraction",
	4: "Constructor Destructor",
}

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{name: name, age: age}
}

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) GetAge() int {
	return p.age
}

func (p *Person) SetAge(age int) {
	if age > 0 {
		p.age = age
	}
}

func Encapsulation() {
	person := NewPerson("John Doe", 30)
	fmt.Println(person.GetName())
	fmt.Println(person.GetAge())

	person.SetAge(35)
	fmt.Println(person.GetAge())
}

type Animal struct {
	name string
}

func (a *Animal) GetName() string {
	return a.name
}

type Dog struct {
	Animal
	breed string
}

func NewDog(name, breed string) *Dog {
	return &Dog{Animal: Animal{name: name}, breed: breed}
}

func Inheritance() {
	dog := NewDog("Buddy", "Golden Retriever")
	fmt.Println(dog.GetName())
	fmt.Println(dog.breed)
}

type Speaker interface {
	Speak() string
}

func (d *Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c *Cat) Speak() string {
	return "Meow!"
}

func Polymorphism() {
	animals := []Speaker{&Dog{}, &Cat{}}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width, height float64
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func Abstraction() {
	var s Shape

	s = &Rectangle{width: 10, height: 5}
	fmt.Println("Rectangle Area:", s.Area())

	s = &Circle{radius: 7}
	fmt.Println("Circle Area:", s.Area())
}

type File struct {
	name string
}

func NewFile(name string) *File {
	fmt.Println("Opening file:", name)
	return &File{name: name}
}

func (f *File) Close() {
	fmt.Println("Closing file:", f.name)
}

func ConstructorDestructor() {
	file := NewFile("test.txt")
	defer file.Close()

	fmt.Println("File operations")
}

func PrintOop() {
	for i, oFunc := range oops {
		fmt.Println("Algorithm name:", oopsName[i])
		oFunc()
	}
}
