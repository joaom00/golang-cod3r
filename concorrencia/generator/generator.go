package main

import (
	"fmt"

	"github.com/joaom00/golang/html"
)

func main() {
	t1 := html.Titulo("https://www.google.com", "https://www.cod3r.com.br")
	t2 := html.Titulo("https://www.amazon.com", "https://www.youtube.com")
	fmt.Println("Primeiros:", <-t1, "|", <-t2)
	fmt.Println("Segundos:", <-t1, "|", <-t2)
}
