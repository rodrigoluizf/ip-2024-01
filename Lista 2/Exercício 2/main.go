package main

import (
	"fmt"
)

func main() {
	var populacaoA, populacaoB int
	var anos int

	fmt.Scan(&populacaoA)
	fmt.Scan(&populacaoB)

	for populacaoA < populacaoB {
		populacaoA += int(float64(populacaoA) * 0.03)
		populacaoB += int(float64(populacaoB) * 0.015)
		anos++
	}

	fmt.Printf("ANOS = %d\n", anos)
}
