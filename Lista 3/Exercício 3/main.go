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

	for i := n - 1; i >= 0; i-- {
		fmt.Printf("%d ", valores[i])
	}
	fmt.Println()
}
