package main

import (
	"fmt"
)

func troca(vetor []int, x, y int) {
	vetor[x], vetor[y] = vetor[y], vetor[x]
}

func saoOpostos(a, b int) bool {
	return a+b == 0
}

func trocaOpostosSeMenor(vetor []int) {
	n := len(vetor)
	for i := 0; i < n/2; i++ {
		j := n - 1 - i
		if i < j && saoOpostos(vetor[i], vetor[j]) && vetor[i] < vetor[j] {
			troca(vetor, i, j)
		}
	}
}

func main() {
	var casos int
	fmt.Scan(&casos)

	for c := 0; c < casos; c++ {
		var n int
		fmt.Scan(&n)

		vetor := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&vetor[i])
		}

		trocaOpostosSeMenor(vetor)

		for i := 0; i < n; i++ {
			fmt.Printf("%d ", vetor[i])
		}
		fmt.Println()
	}
}