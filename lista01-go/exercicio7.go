package main

import (
	"fmt"
	"math"
)

func main() {
	var p, m, f, c float64

	fmt.Print("Digite a temperatura em Fahrenheit: ")
	fmt.Scan(&f)
	fmt.Print("Digite a quantidade de chuva em polegadas: ")
	fmt.Scan(&p)

	c = 5 * (f - 32) / 9
	m = 25.4 * p

	fmt.Printf("O VALOR EM CELSIUS= %.2f\n", math.Round(c*100)/100)
	fmt.Printf("A QUANTIDADE DE CHUVA É= %.2f\n", math.Round(m*100)/100)
}
