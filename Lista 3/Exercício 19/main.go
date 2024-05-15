package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	var numAtual, numAnterior int

	primeiroNumero := true

	for i := 0; i < N; i++ {
		fmt.Scan(&numAtual)

		if primeiroNumero {
			fmt.Println(numAtual)
			primeiroNumero = false
		} else if numAtual != numAnterior {

			fmt.Println(numAtual)
		}
		numAnterior = numAtual
	}
}
