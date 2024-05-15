package main

import (
	"fmt"
)

func main() {
	var n, numeroAtual, numeroAnterior, comprimentoMax, comprimentoAtual int
	fmt.Scan(&n)

	if n > 0 {
		comprimentoMax = 0
		comprimentoAtual = 0

		for i := 0; i < n; i++ {
			fmt.Scan(&numeroAtual)
			if i == 0 {
				numeroAnterior = numeroAtual
				continue
			}
			if numeroAtual > numeroAnterior {
				comprimentoAtual++
				if comprimentoAtual > comprimentoMax {
					comprimentoMax = comprimentoAtual
				}
			} else {
				comprimentoAtual = 0
			}
			numeroAnterior = numeroAtual
		}
	}

	fmt.Printf("O comprimento do segmento crescente maximo e: %d\n", comprimentoMax)
}
