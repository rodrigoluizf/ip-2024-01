package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	V := make([]int, N)
	for i := range V {
		fmt.Scan(&V[i])
	}

	W := make([]int, N)
	for i, v := range V {
		W[N-1-i] = v
	}

	for _, v := range V {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("\n")

	for _, w := range W {
		fmt.Printf("%d ", w)
	}
	fmt.Printf("\n")

	maiorV, menorW := V[0], W[0]
	for _, v := range V {
		if v > maiorV {
			maiorV = v
		}
	}
	for _, w := range W {
		if w < menorW {
			menorW = w
		}
	}

	fmt.Printf("%d\n%d\n", maiorV, menorW)
}
