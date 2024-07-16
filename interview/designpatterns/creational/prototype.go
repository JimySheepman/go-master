package creational

import "fmt"

type Cloneable interface {
	Clone() Cloneable
}

type Person struct {
	name string
	age  int
}

func (p *Person) Clone() Cloneable {
	return &Person{name: p.name, age: p.age}
}

func Prototype() {
	p1 := &Person{name: "John", age: 30}
	p2 := p1.Clone().(*Person)

	fmt.Println(p1)
	fmt.Println(p2)

	p2.name = "Doe"
	fmt.Println(p1)
	fmt.Println(p2)
}
