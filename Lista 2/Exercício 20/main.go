package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	soma := 0
	divisores := ""

	for i := 1; i < n; i++ {
		if n%i == 0 {
			soma += i
			divisores += fmt.Sprintf("%d + ", i)
		}
	}

	divisores = divisores[:len(divisores)-2]

	mensagem := "NUMERO NAO E PERFEITO"
	if soma == n {
		mensagem = "NUMERO PERFEITO"
	}

	fmt.Printf("%d = %s = %d (%s)\n", n, divisores, soma, mensagem)
}
