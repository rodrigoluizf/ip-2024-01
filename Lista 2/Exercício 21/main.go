package main

import (
	"fmt"
)

func main() {
	var tamanho int
	for {
		fmt.Scan(&tamanho)
		if tamanho == 0 {
			break
		}

		sequencia := make([]float64, tamanho)
		for i := range sequencia {
			fmt.Scan(&sequencia[i])
		}

		ordenada := true
		for i := 0; i < tamanho-1; i++ {
			if sequencia[i] > sequencia[i+1] {
				ordenada = false
				break
			}
		}

		if ordenada {
			fmt.Println("ORDENADA")
		} else {
			fmt.Println("DESORDENADA")
		}
	}
}
