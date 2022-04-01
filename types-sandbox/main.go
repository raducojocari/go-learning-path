package main

import "fmt"

type Noisy interface {
	MakeNoise() string
}

type Animal interface {
	Noisy
}

func (a *Animal) MakeNoise() string {
	return "RAAGGH"
}

type Wolf Animal

func (w *Wolf) MakeNoise() string {
	return fmt.Sprintf("The %s goes AUUUU", w.name)
}

func NewWolf() Wolf {
	return Wolf{
		name: "WOLF",
	}
}

func main() {
	w := NewWolf()
	println(w.MakeNoise())

	var a Animal = NewWolf()
	println(a.MakeNoise())
}
