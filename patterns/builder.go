package main

import "fmt"

// Паттерн "Строитель" используется для задания значений в структуре данных пользователем, при этом методы, реализующие
// задание полей структуры возвращают интерфейс взаимодействия для возможности каскадного вызова методов.
// Используется для конфигурации логера Zerolog (zerolog.New(loggerFile).With().Timestamp().Caller().Logger())

type CarBuilder interface {
	SetWheels(wheels int) CarBuilder
	SetSeats(seats int) CarBuilder
	SetModel(structure string) CarBuilder

	Build() Car
}

type Car struct {
	Wheels int
	Seats  int
	Model  string
}

type carBuilder struct {
	car Car
}

func NewConcreteCarBuilder() CarBuilder {
	return &carBuilder{}
}

func (b *carBuilder) SetWheels(wheels int) CarBuilder {
	b.car.Wheels = wheels
	return b
}

func (b *carBuilder) SetSeats(seats int) CarBuilder {
	b.car.Seats = seats
	return b
}

func (b *carBuilder) SetModel(mark string) CarBuilder {
	b.car.Model = mark
	return b
}

func (b *carBuilder) Build() Car {
	return b.car
}

func main() {
	builder := NewConcreteCarBuilder()
	car := builder.SetWheels(4).SetSeats(5).SetModel("ИКАРУС 250").Build()
	fmt.Printf("%+v", car)
}
