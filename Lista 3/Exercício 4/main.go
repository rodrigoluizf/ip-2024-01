package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	if n >= 5000 {
		fmt.Println("O número de elementos deve ser menor que 5000.")
		return
	}

	valores := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&valores[i])
	}

	for _, valor := range valores {
		fmt.Printf("%d ", valor)
	}
	fmt.Println()
}
