package main

import (
	"fmt"
)

// Estrutura da pilha
type Stack struct {
	elementos []rune
}

// Push: adiciona um item no topo
func (s *Stack) Push(r rune) {
	s.elementos = append(s.elementos, r)
}

// Pop: remove o item do topo e retorna
func (s *Stack) Pop() (rune, bool) {
	if len(s.elementos) == 0 {
		return 0, false
	}
	top := len(s.elementos) - 1
	valor := s.elementos[top]
	s.elementos = s.elementos[:top]
	return valor, true
}

// Verifica se uma palavra é palíndromo com pilha
func ehPalindromo(palavra string) bool {
	var pilha Stack

	for _, letra := range palavra {
		pilha.Push(letra)
	}

	invertida := ""
	for range palavra {
		letra, _ := pilha.Pop()
		invertida += string(letra)
	}

	return palavra == invertida
}

func main() {
	testes := []string{"radar", "arara", "golang", "level", "hello"}

	for _, palavra := range testes {
		if ehPalindromo(palavra) {
			fmt.Printf("\"%s\" é palíndromo\n", palavra)
		} else {
			fmt.Printf("\"%s\" não é palíndromo\n", palavra)
		}
	}
}
