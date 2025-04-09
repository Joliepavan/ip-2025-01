package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m float64

	fmt.Print("Digite o salário: ")
	fmt.Scan(&n)

	if n <= 300 {
		m = n + 0.5*n
	} else {
		m = n + 0.3*n
	}

	fmt.Printf("SALÁRIO COM REAJUSTE= %.2f\n", math.Round(m*100)/100)
}
