package main

import (
	"fmt"
)

func ParOuImpar(n int) string {
	if n%2 == 0 {
		return "par"
	}
	return "Ímpar"
}

func main() {
	var n int
	fmt.Scan(&n)
	
	fmt.Println(ParOuImpar(n))
}
