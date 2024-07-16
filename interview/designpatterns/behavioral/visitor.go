package behavioral

import "fmt"

type Element interface {
	Accept(visitor Visitorer)
}

type Visitorer interface {
	VisitConcreteElementA(element *ConcreteElementA)
	VisitConcreteElementB(element *ConcreteElementB)
}

type ConcreteElementA struct{}

func (e *ConcreteElementA) Accept(visitor Visitorer) {
	visitor.VisitConcreteElementA(e)
}

func (e *ConcreteElementA) OperationA() {
	fmt.Println("ConcreteElementA operation")
}

type ConcreteElementB struct{}

func (e *ConcreteElementB) Accept(visitor Visitorer) {
	visitor.VisitConcreteElementB(e)
}

func (e *ConcreteElementB) OperationB() {
	fmt.Println("ConcreteElementB operation")
}

type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Println("Visitor visiting ConcreteElementA")
	element.OperationA()
}

func (v *ConcreteVisitor) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Println("Visitor visiting ConcreteElementB")
	element.OperationB()
}

func Visitor() {
	elements := []Element{&ConcreteElementA{}, &ConcreteElementB{}}
	visitor := &ConcreteVisitor{}

	for _, element := range elements {
		element.Accept(visitor)
	}
}
