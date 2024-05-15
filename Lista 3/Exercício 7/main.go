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
		for _, valor := range vetor {
			if valor > maiorValor {
				maiorValor = valor
			}
		}

		for i := 0; i <= maiorValor; i++ {
			frequencia := 0
			for _, valor := range vetor {
				if valor <= i {
					frequencia++
				}
			}
			fmt.Printf("(%d) %d\n", i, frequencia)
		}
	}
}
