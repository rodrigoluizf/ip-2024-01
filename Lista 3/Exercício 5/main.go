package main

import (
	"fmt"
)

func main() {
	var N int

	for {
		fmt.Scan(&N)
		if N == 0 {
			break
		}

		vetor := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&vetor[i])
		}

		maiorValor := vetor[0]
		indiceMaiorValor := 0

		for i, valor := range vetor {
			if valor > maiorValor {
				maiorValor = valor
				indiceMaiorValor = i
			}
		}

		fmt.Printf("%d %d\n", indiceMaiorValor, maiorValor)
	}
}
