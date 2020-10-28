package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Jose Joao":    1231.31,
		"Maria Silva":  1231.12,
		"Pedro Junior": 15123.31,
	}

	funcsESalarios["Rafael Filho"] = 1350.0
	delete(funcsESalarios, "Inexistente")
	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
