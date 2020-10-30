package main

import (
	"fmt"
	"time"

	"github.com/dmorais8/html"
)

func oMaisRapido(url1, url2, url3 string) string {
	c1 := html.Titulos(url1)
	c2 := html.Titulos(url2)
	c3 := html.Titulos(url3)

	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam"
	default:
		return "Sem resposta ainda"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://www.youtube.com",
		"https://www.amazon.com",
		"https://www.google.com",
	)

	fmt.Println(campeao)
}
