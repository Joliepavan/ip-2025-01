package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, delta float64

	fmt.Print("Insira a: ")
	fmt.Scan(&a)
	fmt.Print("Insira b: ")
	fmt.Scan(&b)
	fmt.Print("Insira c: ")
	fmt.Scan(&c)

	delta = math.Pow(b, 2.0) - 4*a*c

	fmt.Printf("O VALOR DE DELTA É = %.2f\n", math.Round(delta*100)/100)
}
