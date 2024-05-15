package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int
	fmt.Scan(&n)

	fatorial := big.NewInt(1)
	for i := 1; i <= n; i++ {
		fatorial.Mul(fatorial, big.NewInt(int64(i)))
	}

	fmt.Printf("%d! = %s\n", n, fatorial.String())
}
