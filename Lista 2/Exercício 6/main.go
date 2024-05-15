package main

import (
	"fmt"
)

func main() {
	var n, numeroAtual, numeroAnterior, maiorSubsequencia, subsequenciaAtual int
	fmt.Scan(&n)

	if n > 0 {
		fmt.Scan(&numeroAnterior)
		maiorSubsequencia = 1
		subsequenciaAtual = 1

		for i := 1; i < n; i++ {
			fmt.Scan(&numeroAtual)
			if numeroAtual == numeroAnterior {
				subsequenciaAtual++
				if subsequenciaAtual > maiorSubsequencia {
					maiorSubsequencia = subsequenciaAtual
				}
			} else {
				subsequenciaAtual = 1
			}
			numeroAnterior = numeroAtual
		}
	}

	fmt.Printf("A maior subsequência de elementos iguais possui %d elementos.\n", maiorSubsequencia)
}
