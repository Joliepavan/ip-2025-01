package main

import (
	"fmt"
)

func main() {
	var n int
	var x float64

	fmt.Print("Digite um número: ")
	fmt.Scan(&n)

	if n%3 == 0 {
		x = float64(n) * 10 / 3
	} else {
		x = (float64(n-n%3) / 3) * 10 + float64(n%3)*5
	}

	fmt.Printf("O VALOR A PAGAR É = %.2f\n", x)
}
