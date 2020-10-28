package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	velocidadeAtual int
	turboLigado     bool
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	carro1 := ferrari{"f40", 0, false}
	carro1.ligarTurbo()

	var carro2 esportivo = &ferrari{"f40", 0, false}
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)
}
