package main

import "fmt"

func main() {
	aprovados := make(map[int]string)

	aprovados[123456789] = "Maria"
	aprovados[123412344] = "Ana"
	aprovados[123234153] = "Pedro"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}
}
