package main

import (
	"fmt"
)

func main() {
	var alunos [3][3]string

	for i := 0; i < 3; i++ {
		fmt.Printf("Informe o índice do aluno %d: ", i+1)
		fmt.Scan(&alunos[i][0])

		fmt.Printf("Informe o RA do aluno %d: ", i+1)
		fmt.Scan(&alunos[i][1])

		fmt.Printf("Informe o nome do aluno %d: ", i+1)
		fmt.Scan(&alunos[i][2])
	}

	fmt.Println("\nResumo dos Alunos")
	fmt.Println("Índice\tRA\tNome")
	for i := 0; i < 3; i++ {
    //"%s\t%s\t%s\n" é como se fosse o tab, da um espaço maior entre os itens
		fmt.Printf("%s\t%s\t%s\n", alunos[i][0], alunos[i][1], alunos[i][2])
	}
}
