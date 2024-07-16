package behavioral

import "fmt"

type Observerer interface {
	Update(message string)
}

type Subject interface {
	Register(observer Observerer)
	Deregister(observer Observerer)
	NotifyAll()
}

// ConcreteSubject yapısı
type ConcreteSubject struct {
	observers []Observerer
	message   string
}

func (s *ConcreteSubject) Register(observer Observerer) {
	s.observers = append(s.observers, observer)
}

func (s *ConcreteSubject) Deregister(observer Observerer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ConcreteSubject) NotifyAll() {
	for _, observer := range s.observers {
		observer.Update(s.message)
	}
}

func (s *ConcreteSubject) UpdateMessage(message string) {
	s.message = message
	s.NotifyAll()
}

// ConcreteObserver yapısı
type ConcreteObserver struct {
	name string
}

func (o *ConcreteObserver) Update(message string) {
	fmt.Println(o.name, "received message:", message)
}

func Observer() {
	subject := &ConcreteSubject{}

	observer1 := &ConcreteObserver{name: "Observer1"}
	observer2 := &ConcreteObserver{name: "Observer2"}

	subject.Register(observer1)
	subject.Register(observer2)

	subject.UpdateMessage("Hello Observers!")

	subject.Deregister(observer1)

	subject.UpdateMessage("Hello Observer2!")
}
