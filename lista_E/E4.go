package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	var v float64
	var S float64
	numeros := make([]float64, 0)

	fmt.Scan(&N)

	if N%2 != 0 {
		return
	}

	for i := 0; i < N; i++ {
		fmt.Scan(&v)
		numeros = append(numeros, v)
	}

	for i := 0; i < N/2; i++ {
		diferenca := numeros[i] - numeros[N-1-i]
		S += math.Pow(diferenca, 3)
	}

	fmt.Printf("%.2f\n", S)
}
