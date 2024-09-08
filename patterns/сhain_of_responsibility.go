package main

import "fmt"

// Паттерн "Цепочка вызовов" позволяет передавать запросы по цепочке вызовов пока не найдется необходимый обработчик

type Handler interface {
	SetNext(handler Handler) Handler
	Handle(request string) string
}

type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(handler Handler) Handler {
	h.next = handler
	return handler
}

func (h *BaseHandler) Handle(request string) string {
	if h.next != nil {
		return h.next.Handle(request)
	}
	return "No handler available"
}

// ConcreteHandlerA обрабатывает запросы определенного типа `A`
type ConcreteHandlerA struct {
	BaseHandler
}

func (h *ConcreteHandlerA) Handle(request string) string {
	if request == "A" {
		return "ConcreteHandlerA handled " + request
	}
	return h.BaseHandler.Handle(request)
}

// ConcreteHandlerB обрабатывает запросы другого типа `B`
type ConcreteHandlerB struct {
	BaseHandler
}

func (h *ConcreteHandlerB) Handle(request string) string {
	if request == "B" {
		return "ConcreteHandlerB handled " + request
	}
	return h.BaseHandler.Handle(request)
}

func main() {
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}

	handlerA.SetNext(handlerB)

	requests := []string{"A", "B", "C"}

	for _, req := range requests {
		fmt.Println("Request:", req, "Result:", handlerA.Handle(req))
	}
}
