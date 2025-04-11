package main

import (
  "fmt"
)

type Produto struct {
  Nome       string
  Preco      float64
  Quantidade int
}

func main() {
  prod := Produto{
    Nome:       "Guaravita",
    Preco:      2.50,
    Quantidade: 10,
  }

  fmt.Println("Produto:", prod.Nome)
  fmt.Printf("Preço: R$ %.2f\n", prod.Preco)
  fmt.Printf("Estoque: %d unidades\n", prod.Quantidade)
}
