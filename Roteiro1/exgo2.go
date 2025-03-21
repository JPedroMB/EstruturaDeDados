package main

import (
	"fmt"
)

func SomaAte(n int) int {
	soma := 0

	for i := 1; i <= n; i++ {
		soma += i
	}
	return soma
}
func main() {
	var n int
	fmt.Scanln(&n)
	resultado := SomaAte(n)
	fmt.Println(resultado)
}
