package main

import "fmt"

// Паттерн "Стратегия" изолирует семейство алгоритмов, обертывая их общим интерфейсом. Каждый алгоритм представляет
// собой свою структуру и обрабатывает вызов по своему.

type NavigationStrategy interface {
	Route() string
}

type RoadStrategy struct{}

func (r *RoadStrategy) Route() string {
	return "Route calculated using road map"
}

type PublicTransportStrategy struct{}

func (p *PublicTransportStrategy) Route() string {
	return "Route calculated using public transport"
}

type Navigator struct {
	strategy NavigationStrategy
}

func (n *Navigator) SetStrategy(strategy NavigationStrategy) {
	n.strategy = strategy
}

func (n *Navigator) Navigate() {
	fmt.Println(n.strategy.Route())
}

func main() {
	navigator := &Navigator{}

	roadStrategy := &RoadStrategy{}
	navigator.SetStrategy(roadStrategy)
	navigator.Navigate()

	publicTransportStrategy := &PublicTransportStrategy{}
	navigator.SetStrategy(publicTransportStrategy)
	navigator.Navigate()
}
