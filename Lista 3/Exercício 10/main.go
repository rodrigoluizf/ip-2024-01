package main

import (
	"fmt"
)

func main() {
	var N, nota int
	fmt.Scan(&N)

	notas := make([]int, N)
	frequencia := make(map[int]int)
	maiorNota := 0
	indiceMaiorNota := -1

	for i := 0; i < N; i++ {
		fmt.Scan(&nota)
		notas[i] = nota
		frequencia[nota]++

		if nota > maiorNota {
			maiorNota = nota
			indiceMaiorNota = i
		}
	}

	ultimaNota := notas[N-1]
	fmt.Printf("Nota %d, %d vezes\n", ultimaNota, frequencia[ultimaNota])
	if indiceMaiorNota != -1 {
		fmt.Printf("Nota %d, indice %d\n", maiorNota, indiceMaiorNota)
	}
}
