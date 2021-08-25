package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Teste " + string(97)) // output => a

	fmt.Println("Teste " + strconv.Itoa(255)) // output => Teste 255

	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}
}
