package main

import (
	"fmt"
)

func main() {
	var matricula int
	var horasTrabalhadas, valorPorHora, salario float64

	for {
		fmt.Scanf("%d %f %f\n", &matricula, &horasTrabalhadas, &valorPorHora)
		if matricula == 0 {
			break
		}

		salario = horasTrabalhadas * valorPorHora
		fmt.Printf("%d %.2f\n", matricula, salario)
	}
}
