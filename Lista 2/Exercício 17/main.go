package main

import (
	"fmt"
)

func main() {
	var codigo, codigoMaiorLucro, codigoMaisVendida uint64
	var precoCompra, precoVenda, lucro, lucroTotal, maiorLucro, totalCompra, totalVenda float64
	var vendas, maisVendida, lucroMenor10, lucroEntre10e20, lucroMaior20 int

	for {
		_, err := fmt.Scan(&codigo, &precoCompra, &precoVenda, &vendas)
		if err != nil {
			break
		}

		lucro = (precoVenda - precoCompra) * float64(vendas)
		percentualLucro := (precoVenda - precoCompra) / precoCompra * 100

		totalCompra += precoCompra * float64(vendas)
		totalVenda += precoVenda * float64(vendas)
		lucroTotal += lucro

		if percentualLucro < 10 {
			lucroMenor10++
		} else if percentualLucro <= 20 {
			lucroEntre10e20++
		} else {
			lucroMaior20++
		}

		if lucro > maiorLucro {
			maiorLucro = lucro
			codigoMaiorLucro = codigo
		}

		if vendas > maisVendida {
			maisVendida = vendas
			codigoMaisVendida = codigo
		}
	}

	percentualLucroTotal := lucroTotal / totalCompra * 100

	fmt.Printf("Quantidade de mercadorias que geraram lucro menor que 10%%: %d\n", lucroMenor10)
	fmt.Printf("Quantidade de mercadorias que geraram lucro maior ou igual a 10%% e menor ou igual a 20%%: %d\n", lucroEntre10e20)
	fmt.Printf("Quantidade de mercadorias que geraram lucro maior do que 20%%: %d\n", lucroMaior20)
	fmt.Printf("Codigo da mercadoria que gerou maior lucro: %d\n", codigoMaiorLucro)
	fmt.Printf("Codigo da mercadoria mais vendida: %d\n", codigoMaisVendida)
	fmt.Printf("Valor total de compras: %.2f, valor total de vendas: %.2f e percentual de lucro total: %.2f%%\n", totalCompra, totalVenda, percentualLucroTotal)
}
