package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N int

	for {
		_, err := fmt.Scan(&N)
		if err != nil {
			break
		}

		binario := strconv.FormatInt(int64(N), 2)
		fmt.Println(binario)
	}
}
