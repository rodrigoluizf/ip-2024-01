package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)

	pontos := make([][3]float64, N)
	for i := range pontos {
		fmt.Scan(&pontos[i][0], &pontos[i][1], &pontos[i][2])
	}

	for i := 0; i < N-1; i++ {
		vetor := [3]float64{
			pontos[i+1][0] - pontos[i][0],
			pontos[i+1][1] - pontos[i][1],
			pontos[i+1][2] - pontos[i][2],
		}

		maxValorAbsoluto := math.Abs(vetor[0])
		for _, valor := range vetor[1:] {
			valorAbsoluto := math.Abs(valor)
			if valorAbsoluto > maxValorAbsoluto {
				maxValorAbsoluto = valorAbsoluto
			}
		}

		fmt.Printf("%.2f\n", maxValorAbsoluto)
	}
}
