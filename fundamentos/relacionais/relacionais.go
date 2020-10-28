package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings:", "banana" == "banana")

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas:", d1 == d2)

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"João"}
	p2 := Pessoa{"João"}
	fmt.Println("Pessoa:", p1 == p2)
}
