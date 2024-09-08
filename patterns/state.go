package main

import "fmt"

// Паттерн "Состояние" позволяет объекту менять свое поведение в зависимости от своего состояния

type State interface {
	Handle(context *TrafficLight)
}

// TrafficLight представляет контекст, в котором будут меняться состояния
type TrafficLight struct {
	State State
}

func (t *TrafficLight) SetState(state State) {
	t.State = state
}

func (t *TrafficLight) Request() {
	t.State.Handle(t)
}

// GreenLightState состояние зелёного света
type GreenLightState struct{}

func (g *GreenLightState) Handle(context *TrafficLight) {
	fmt.Println("Green light - Go!")
	context.SetState(&YellowLightState{})
}

// YellowLightState состояние жёлтого света
type YellowLightState struct{}

func (y *YellowLightState) Handle(context *TrafficLight) {
	fmt.Println("Yellow light - Prepare to stop")
	context.SetState(&RedLightState{})
}

// RedLightState состояние красного света
type RedLightState struct{}

func (r *RedLightState) Handle(context *TrafficLight) {
	fmt.Println("Red light - Stop")
	context.SetState(&GreenLightState{})
}

func main() {
	light := &TrafficLight{State: &GreenLightState{}}
	light.Request()
	light.Request()
	light.Request()
	light.Request()
}
