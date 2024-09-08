package main

import "fmt"

// Паттерн "Команда" позволяет нескольким сущностям взаимодействовать с одним, изменяя только его некоторые свойства,
// скрывая реализацию всего объекта

type Command interface {
	Execute()
}

type Light struct {
	IsOn bool
}

type TurnOnCommand struct {
	Light *Light
}

func (c *TurnOnCommand) Execute() {
	c.Light.IsOn = true
	fmt.Println("Light is turned on")
}

type TurnOffCommand struct {
	Light *Light
}

func (c *TurnOffCommand) Execute() {
	c.Light.IsOn = false
	fmt.Println("Light is turned off")
}

func main() {
	light := &Light{}

	var turnOn Command = &TurnOnCommand{Light: light}
	var turnOff Command = &TurnOffCommand{Light: light}

	turnOn.Execute()
	turnOff.Execute()
}
