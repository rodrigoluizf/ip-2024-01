package main

import (
	"fmt"
)

func main() {
	var n, i, K int
	var s, B, R float64

	fmt.Scan(&n, &i, &K, &s)

	fmt.Println("Tabuada de soma:")
	B = float64(i)
	for k := 0; k < K; k++ {
		R = float64(n) + B
		fmt.Printf("%.2f + %.2f = %.2f\n", float64(n), B, R)
		B += s
	}

	fmt.Println("Tabuada de subtracao:")
	B = float64(i)
	for k := 0; k < K; k++ {
		R = float64(n) - B
		fmt.Printf("%.2f - %.2f = %.2f\n", float64(n), B, R)
		B += s
	}

	fmt.Println("Tabuada de multiplicacao:")
	B = float64(i)
	for k := 0; k < K; k++ {
		R = float64(n) * B
		fmt.Printf("%.2f x %.2f = %.2f\n", float64(n), B, R)
		B += s
	}

	fmt.Println("Tabuada de divisao:")
	B = float64(i)
	for k := 0; k < K; k++ {
		if B != 0 {
			R = float64(n) / B
			fmt.Printf("%.2f / %.2f = %.2f\n", float64(n), B, R)
		} else {
			fmt.Printf("%.2f / %.2f = Infinito\n", float64(n), B)
		}
		B += s
	}
}
