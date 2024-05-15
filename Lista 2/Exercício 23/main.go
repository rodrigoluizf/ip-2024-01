package main

import (
	"fmt"
)

func somaDivisores(num int) int {
	soma := 0
	for i := 1; i < num; i++ {
		if num%i == 0 {
			soma += i
		}
	}
	return soma
}

func main() {
	var n int
	fmt.Scan(&n)

	encontrados := 0
	for a := 2; encontrados < n; a++ {
		b := somaDivisores(a)
		if b > a && somaDivisores(b) == a {
			fmt.Printf("(%d,%d)\n", a, b)
			encontrados++
		}
	}
}
