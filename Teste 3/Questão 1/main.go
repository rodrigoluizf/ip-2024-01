package main

import (
	"fmt"
	"strings"
)

func inverterPalavra(palavra string) string {
	runes := []rune(palavra)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	var linha string
	for {
		_, err := fmt.Scan(&linha)
		if err != nil {
			break 
		}

		palavras := strings.Split(linha, " ")
		for _, palavra := range palavras {
			fmt.Print(inverterPalavra(palavra), " ")
		}
		fmt.Println()
	}
}