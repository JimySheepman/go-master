package behavioral

import "fmt"

type Mementor struct {
	state string
}

type Originator struct {
	state string
}

func (o *Originator) SetState(state string) {
	o.state = state
}

func (o *Originator) SaveStateToMemento() *Mementor {
	return &Mementor{state: o.state}
}

func (o *Originator) GetStateFromMemento(memento *Mementor) {
	o.state = memento.state
}

func (o *Originator) GetState() string {
	return o.state
}

type Caretaker struct {
	mementoList []*Mementor
}

func (c *Caretaker) Add(memento *Mementor) {
	c.mementoList = append(c.mementoList, memento)
}

func (c *Caretaker) Get(index int) *Mementor {
	return c.mementoList[index]
}

func Memento() {
	originator := &Originator{}
	caretaker := &Caretaker{}

	originator.SetState("State1")
	originator.SetState("State2")
	caretaker.Add(originator.SaveStateToMemento())

	originator.SetState("State3")
	caretaker.Add(originator.SaveStateToMemento())

	originator.SetState("State4")
	fmt.Println("Current State:", originator.GetState())

	originator.GetStateFromMemento(caretaker.Get(0))
	fmt.Println("First saved State:", originator.GetState())

	originator.GetStateFromMemento(caretaker.Get(1))
	fmt.Println("Second saved State:", originator.GetState())
}
