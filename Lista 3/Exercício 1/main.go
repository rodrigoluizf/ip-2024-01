package main

import (
	"fmt"
)

func main() {
	var N, M, num int
	fmt.Scan(&N)

	vetorV := make(map[int]bool, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&num)
		vetorV[num] = true
	}

	fmt.Scan(&M)
	for i := 0; i < M; i++ {
		fmt.Scan(&num)
		if _, ok := vetorV[num]; ok {
			fmt.Println("ACHEI")
		} else {
			fmt.Println("NAO ACHEI")
		}
	}
}