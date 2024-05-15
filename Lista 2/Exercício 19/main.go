package main

import (
	"fmt"
)

func main() {
	var m int
	fmt.Scan(&m)

	for n := 1; n <= m; n++ {
		start := n*n - n + 1
		fmt.Printf("%d*%d*%d =", n, n, n)
		for i := 0; i < n; i++ {
			fmt.Printf(" %d", start+2*i)
			if i < n-1 {
				fmt.Print("+")
			}
		}
		fmt.Println()
	}
}
