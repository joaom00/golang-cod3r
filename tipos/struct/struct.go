package main

import "fmt"

type Produto struct {
	nome     string
	preco    float64
	desconto float64
}

func (p Produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	produto1 := Produto{
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}
	fmt.Println(produto1, produto1.precoComDesconto())
}
