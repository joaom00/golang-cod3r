package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type eu struct {
	ID       string
	Name     string
	LastName string
	Age      int
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
}

type bwm7 struct{}

func (b bwm7) ligarTurbo() {
	fmt.Println("Turbo...")
}

func (b bwm7) fazerBaliza() {
	fmt.Println("Baliza...")
}

func main() {
	var b esportivoLuxuoso = bwm7{}
	b.ligarTurbo()
	b.fazerBaliza()
}
