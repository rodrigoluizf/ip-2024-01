package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)

	numeros := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&numeros[i])
	}

	sort.Ints(numeros)

	var mediana float64
	if N%2 == 0 {
		mediana = (float64(numeros[N/2-1]) + float64(numeros[N/2])) / 2
	} else {
		mediana = float64(numeros[N/2])
	}

	fmt.Printf("%.2f\n", mediana)
}
