package main

import (
	"fmt"
)

func main() {
	var valorIngresso, valorInicial, valorFinal float64
	fmt.Scan(&valorIngresso, &valorInicial, &valorFinal)

	if valorInicial >= valorFinal {
		fmt.Println("INTERVALO INVALIDO.")
		return
	}

	var melhorValorFinal, melhorLucro float64
	var numeroIngressos int
	for v := valorInicial; v <= valorFinal; v++ {
		var n int
		if v < valorIngresso {
			n = 120 + 25*int((valorIngresso-v)/0.5)
		} else if v > valorIngresso {
			n = 120 - 30*int((v-valorIngresso)/0.5)
		} else {
			n = 120
		}

		despesas := 200.0 + 0.05*float64(n)
		lucro := float64(n)*v - despesas

		fmt.Printf("V: %.2f, N: %d, L: %.2f\n", v, n, lucro)

		if lucro > melhorLucro {
			melhorLucro = lucro
			melhorValorFinal = v
			numeroIngressos = n
		}
	}

	if melhorLucro <= 0 {
		fmt.Println("Melhor valor final: 0.00")
		fmt.Println("Lucro: 0.00")
		fmt.Println("Numero de ingressos: 0")
	} else {
		fmt.Printf("Melhor valor final: %.2f\n", melhorValorFinal)
		fmt.Printf("Lucro: %.2f\n", melhorLucro)
		fmt.Printf("Numero de ingressos: %d\n", numeroIngressos)
	}
}
