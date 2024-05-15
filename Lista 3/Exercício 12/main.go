package main

import (
	"fmt"
)

func main() {
	var N, num int
	fmt.Scan(&N)

	frequencia := make(map[int]int)
	var maisFrequente, maiorFrequencia int

	for i := 0; i < N; i++ {
		fmt.Scan(&num)
		frequencia[num]++
		if frequencia[num] > maiorFrequencia || (frequencia[num] == maiorFrequencia && num < maisFrequente) {
			maiorFrequencia = frequencia[num]
			maisFrequente = num
		}
	}

	fmt.Println(maisFrequente)
	fmt.Println(maiorFrequencia)
}