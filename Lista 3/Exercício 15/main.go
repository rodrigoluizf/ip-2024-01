package main

import (
	"fmt"
	"sort"
)

func main() {
	var T, N int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&N)
		numeros := make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Scan(&numeros[j])
		}

		sort.Ints(numeros)

		menorDistancia := 1001
		for k := 1; k < N; k++ {
			distancia := numeros[k] - numeros[k-1]
			if distancia < menorDistancia {
				menorDistancia = distancia
			}
		}

		comparacoes := N * (N - 1) / 2
		fmt.Println(menorDistancia, comparacoes)
	}
}
