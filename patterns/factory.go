package main

import "fmt"

// Паттерн "Фабрика" изолирует создание объекта, позволяет уменьшить дублирование кода и прост в расширении

type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getPower() int {
	return g.power
}

func newAk47() IGun {
	return &Gun{
		name:  "AK47 gun",
		power: 4,
	}
}

func newMusket() IGun {
	return &Gun{
		name:  "Musket gun",
		power: 1,
	}
}

func getGun(gunType string) (IGun, error) {
	switch gunType {
	case "ak47":
		return newAk47(), nil
	case "musket":
		return newMusket(), nil
	default:
		return nil, fmt.Errorf("Wrong gun type passed")
	}
}

func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	fmt.Printf("Gun: %s, Power: %d\n", ak47.getName(), ak47.getPower())
	fmt.Printf("Gun: %s, Power: %d", musket.getName(), musket.getPower())
}
