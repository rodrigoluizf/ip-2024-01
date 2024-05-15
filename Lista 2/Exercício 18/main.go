package main

import (
	"fmt"
)

func mdc(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func mmc(a, b int) int {
	return a * b / mdc(a, b)
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	mmcAB := mmc(a, b)
	mmcFinal := mmc(mmcAB, c)

	fmt.Printf("%5d %5d %5d :%d\n", a, b, c, mdc(a, b))
	fmt.Printf("%5d %5d %5d :%d\n", a/mdc(a, b), b/mdc(a, b), c, mdc(mmcAB, c))
	fmt.Printf("%5d %5d %5d :%d\n", a/mdc(a, b), b/mdc(a, b), c/mdc(mmcAB, c), mmcFinal)
	fmt.Printf("MMC: %d\n", mmcFinal)
}
