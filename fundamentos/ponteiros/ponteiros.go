package main

import "fmt"

func main() {
	i := 1

	var p *int = nil

	p = &i
	*p++
	i++
	fmt.Println(i, *p)
	// go nao tem artimetica de ponteiro
}
