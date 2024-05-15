package main

import (
	"fmt"
	"math"
)

func mdc(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var n float64
	fmt.Scan(&n)

	denominador := int(math.Pow(10, float64(len(fmt.Sprintf("%.0f", n))-2)))

	numerador := int(n * float64(denominador))

	divisor := mdc(numerador, denominador)
	numerador /= divisor
	denominador /= divisor

	fmt.Printf("%d/%d\n", numerador, denominador)
}
