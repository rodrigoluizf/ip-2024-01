package main

import (
	"fmt"
	"math"
)

func main() {
	var n, e float64
	fmt.Scan(&n, &e)

	r := 1.0
	for math.Abs(n-(r*r)) > e {
		r = (r + n/r) / 2
		fmt.Printf("r: %.9f, erro: %.9f\n", r, math.Abs(n-(r*r)))
	}
}
