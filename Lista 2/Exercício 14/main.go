package main

import (
	"fmt"
)

func main() {
	var m, n int
	fmt.Scan(&m, &n)

	for i := 2; i <= m; i++ {
		for j := 1; j < i && j <= n; j++ {
			fmt.Printf("(%d,%d)", i, j)
			if j < i-1 && j < n {
				fmt.Print("-")
			}
		}
		fmt.Println()
	}
}
