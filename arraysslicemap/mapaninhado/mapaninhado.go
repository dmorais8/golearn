package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 1231.21,
			"Guga Pereira":   1233.31,
		},
		"J": {
			"Jose Joao": 1233.11,
		},
		"P": {
			"Pedro Junior": 1241.41,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}

}
