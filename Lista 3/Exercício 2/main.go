package main

import (
	"fmt"
)

func main() {
	var N, K, count int
	for {
		fmt.Scan(&N)
		if N >= 1 && N <= 1000 {
			break
		}
	}

	V := make([]int, N)
	for i := range V {
		fmt.Scan(&V[i])
	}

	fmt.Scan(&K)
	for _, value := range V {
		if value >= K {
			count++
		}
	}

	fmt.Println(count)
}
