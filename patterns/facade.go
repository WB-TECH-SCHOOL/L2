package main

import "fmt"

// Паттерн `Фасад` используется для управления множеством объектов единой структурой для облегчения взаимодействия.
// Представляется некая обертка, помогающая упростить взаимодействие и/или взаимодействовать сразу с группой объектов.
// Используется в структурах сервисного и хендлерного слоя чистой архитектуры:
// type requestService struct {
//	requestRepo    repository.Requests
//	teamRepo       repository.Teams
//	userRepo       repository.Users
//	converter      converters.RequestConverter
//	dbResponseTime time.Duration
//	logger         zerolog.Logger
//}

type Lights struct{}

func (l Lights) TurnOn() {
	fmt.Println("Lights are turned on")
}

func (l Lights) TurnOff() {
	fmt.Println("Lights are turned off")
}

type TV struct{}

func (t TV) TurnOn() {
	fmt.Println("TV is turned on")
}

func (t TV) TurnOff() {
	fmt.Println("TV is turned off")
}

type AirConditioner struct{}

func (a AirConditioner) TurnOn() {
	fmt.Println("Air conditioner is turned on")
}

func (a AirConditioner) TurnOff() {
	fmt.Println("Air conditioner is turned off")
}

type HomeFacade struct {
	lights         Lights
	tv             TV
	airConditioner AirConditioner
}

func NewHomeFacade() HomeFacade {
	return HomeFacade{
		lights:         Lights{},
		tv:             TV{},
		airConditioner: AirConditioner{},
	}
}

func (f HomeFacade) TurnEverythingOn() {
	f.lights.TurnOn()
	f.tv.TurnOn()
	f.airConditioner.TurnOn()
}

func (f HomeFacade) TurnEverythingOff() {
	f.lights.TurnOff()
	f.tv.TurnOff()
	f.airConditioner.TurnOff()
}

func main() {
	home := NewHomeFacade()
	// Включить все устройства
	home.TurnEverythingOn()
	// Выключить все устройства
	home.TurnEverythingOff()
}
