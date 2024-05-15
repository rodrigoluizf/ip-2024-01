package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	elementos := make(map[int]int)
	var valor, unicos int

	for i := 0; i < n; i++ {
		fmt.Scan(&valor)
		elementos[valor]++
	}

	for _, freq := range elementos {
		if freq == 1 {
			unicos++
		}
	}

	fmt.Println(unicos)
}
