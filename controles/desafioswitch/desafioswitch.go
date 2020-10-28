package main

import "fmt"

func notaParaConceito(n float64) string {

	var nota = int(n)
	switch {
	case nota >= 9 && nota <= 10:
		return "A"
	case nota >= 7 && nota < 9:
		return "B"
	case nota >= 5 && nota < 7:
		return "C"
	case nota >= 3 && nota < 5:
		return "D"
	case nota >= 0 && nota < 3:
		return "E"
	default:
		return "Nota inválida"
	}

}

func main() {
	fmt.Println(notaParaConceito(2.1))
}
