package main

import (
	"fmt"
	"math"
)

func main() {
	var n, i int
	var soma float64 = 0.0

	fmt.Print("Digite um número: ")
	fmt.Scan(&n)

	if n <= 1 {
		fmt.Println("Número inválido!")
	} else {
		for i = 1; i <= n; i++ {
			soma += 1.0 / float64(i)
		}
		fmt.Printf("%.6f\n", math.Round(soma*1000000)/1000000)
	}
}

