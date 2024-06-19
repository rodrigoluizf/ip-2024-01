package main

import "fmt"

func pot(x, n int) int {
	if n == 0 {
		return 1
	}
	return x * pot(x, n-1)
}

func main() {
	x := 2
	n := 3
	resultado := pot(x, n)
	fmt.Printf("%d^%d = %d\n", x, n, resultado)
}