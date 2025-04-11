package main

import "fmt"

// Nó da lista
type Node struct {
  data string
  next *Node
}

// Lista encadeada
type LinkedList struct {
  head *Node
}

// Adiciona um novo item ao final
func (l *LinkedList) Append(data string) {
  n := &Node{data: data}
  if l.head == nil {
    l.head = n
    return
  }
  atual := l.head
  for atual.next != nil {
    atual = atual.next
  }
  atual.next = n
}

// Remove o primeiro item da lista
func (l *LinkedList) RemoveFirst() {
  if l.head != nil {
    l.head = l.head.next
  }
}

// Exibe os dados da lista
func (l *LinkedList) Print() {
  atual := l.head
  for atual != nil {
    fmt.Print(atual.data, " -> ")
    atual = atual.next
  }
  fmt.Println("nil")
}

func main() {
  list := LinkedList{}
  list.Append("Assistir melhores momentos gabigol 2019")
  list.Append("Igreja")
  list.Append("Passar raiva no valorant")

  fmt.Println("Lista encadeada:")
  list.Print()

  list.RemoveFirst()

  fmt.Println("Depois de remover o primeiro:")
  list.Print()
}
