package main

import "fmt"

func soma(array []float64, n int) float64 {
	if n == 0 {
		return 0
	}
	return array[n-1] + soma(array, n-1)
}

func main() {
	valores := []float64{1.5, 2.3, 3.7, 4.2}
	resultado := soma(valores, len(valores))
	fmt.Printf("Soma dos valores: %.2f\n", resultado)
}