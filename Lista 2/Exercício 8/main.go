package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	if N < 2 {
		fmt.Println("Campeonato invalido!")
		return
	}

	var contadorFinal int = 1
	for i := 1; i <= N; i++ {
		for j := i + 1; j <= N; j++ {
			fmt.Printf("Final %d: Time%d X Time%d\n", contadorFinal, i, j)
			contadorFinal++
		}
	}
}
