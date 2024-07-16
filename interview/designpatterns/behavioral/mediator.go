package behavioral

import "fmt"

type Mediatorer interface {
	SendMessage(message string, colleague Colleague)
}

type Colleague interface {
	SetMediator(mediator Mediatorer)
	Send(message string)
	Receive(message string)
}

type ConcreteMediator struct {
	colleague1 *ConcreteColleague1
	colleague2 *ConcreteColleague2
}

func (m *ConcreteMediator) SendMessage(message string, colleague Colleague) {
	if colleague == m.colleague1 {
		m.colleague2.Receive(message)
	} else {
		m.colleague1.Receive(message)
	}
}

type ConcreteColleague1 struct {
	mediator Mediatorer
}

func (c *ConcreteColleague1) SetMediator(mediator Mediatorer) {
	c.mediator = mediator
}

func (c *ConcreteColleague1) Send(message string) {
	fmt.Println("Colleague1 sends message:", message)
	c.mediator.SendMessage(message, c)
}

func (c *ConcreteColleague1) Receive(message string) {
	fmt.Println("Colleague1 receives message:", message)
}

type ConcreteColleague2 struct {
	mediator Mediatorer
}

func (c *ConcreteColleague2) SetMediator(mediator Mediatorer) {
	c.mediator = mediator
}

func (c *ConcreteColleague2) Send(message string) {
	fmt.Println("Colleague2 sends message:", message)
	c.mediator.SendMessage(message, c)
}

func (c *ConcreteColleague2) Receive(message string) {
	fmt.Println("Colleague2 receives message:", message)
}

func Mediator() {
	mediator := &ConcreteMediator{}

	colleague1 := &ConcreteColleague1{}
	colleague2 := &ConcreteColleague2{}

	colleague1.SetMediator(mediator)
	colleague2.SetMediator(mediator)

	mediator.colleague1 = colleague1
	mediator.colleague2 = colleague2

	colleague1.Send("Hello")
	colleague2.Send("Hi there")
}
