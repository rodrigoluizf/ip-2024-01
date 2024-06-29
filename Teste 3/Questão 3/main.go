package main

import (
	"fmt"
	"sort"
)

func main() {
	vetor := []int{4, 7, 9, 1, 8}

	sort.Ints(vetor)

	for _, num := range vetor {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}