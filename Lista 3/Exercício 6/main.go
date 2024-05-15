package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	valores := make([]int, n)
	soma := 0

	for i := 0; i < n; i++ {
		fmt.Scan(&valores[i])
		soma += valores[i]
	}

	fmt.Println(soma)
}
