package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)

	if N <= 0 {
		fmt.Println("Numero invalido!")
		return
	}

	isPrime := true
	if N == 1 {
		isPrime = false
	} else {
		for i := 2; i <= int(math.Sqrt(float64(N))); i++ {
			if N%i == 0 {
				isPrime = false
				break
			}
		}
	}

	if isPrime {
		fmt.Println("PRIMO")
	} else {
		fmt.Println("NAO PRIMO")
	}
}
