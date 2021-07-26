package main

import "fmt"

type Item struct {
	produtoID int
	qtde      int
	preco     float64
}

type Pedido struct {
	userID int
	itens  []Item
}

func (p Pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}
	return total
}

func main() {
	pedido := Pedido{
		userID: 1,
		itens: []Item{
			Item{1, 2, 12.10},
			Item{2, 1, 23.49},
			Item{11, 100, 3.13},
		},
	}

	fmt.Printf("Valor total do pedido é R$%.2f", pedido.valorTotal())
}
