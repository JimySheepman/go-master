package behavioral

import "fmt"

type Handler interface {
	SetNext(handler Handler)
	Handle(request string)
}

type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(handler Handler) {
	h.next = handler
}

func (h *BaseHandler) Handle(request string) {
	if h.next != nil {
		h.next.Handle(request)
	}
}

type ConcreteHandler1 struct {
	BaseHandler
}

func (h *ConcreteHandler1) Handle(request string) {
	if request == "request1" {
		fmt.Println("ConcreteHandler1 handled the request")
	} else {
		h.BaseHandler.Handle(request)
	}
}

type ConcreteHandler2 struct {
	BaseHandler
}

func (h *ConcreteHandler2) Handle(request string) {
	if request == "request2" {
		fmt.Println("ConcreteHandler2 handled the request")
	} else {
		h.BaseHandler.Handle(request)
	}
}

func ChainOfResponsibility() {
	handler1 := &ConcreteHandler1{}
	handler2 := &ConcreteHandler2{}

	handler1.SetNext(handler2)

	handler1.Handle("request1")
	handler1.Handle("request2")
	handler1.Handle("request3")
}
