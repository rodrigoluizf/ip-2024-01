package main

import (
	"fmt"
)

func main() {
	var numero, somaPar, somaImpar, contPar, contImpar int
	var mediaPar, mediaImpar float64

	for {
		fmt.Scan(&numero)
		if numero == 0 {
			break
		}

		if numero%2 == 0 {
			somaPar += numero
			contPar++
		} else {
			somaImpar += numero
			contImpar++
		}
	}

	if contPar > 0 {
		mediaPar = float64(somaPar) / float64(contPar)
	}
	if contImpar > 0 {
		mediaImpar = float64(somaImpar) / float64(contImpar)
	}

	fmt.Printf("MEDIA PAR: %.2f\n", mediaPar)
	fmt.Printf("MEDIA IMPAR: %.2f\n", mediaImpar)
}
