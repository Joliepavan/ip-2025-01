package main

import (
	"fmt"
	"math"
)

func main() {
	var h, a, v, ab float64

	fmt.Print("Digite a altura da pirâmide: ")
	fmt.Scan(&h)
	fmt.Print("Digite a aresta da base: ")
	fmt.Scan(&a)

	ab = 3 * (math.Pow(a, 2.0)) * (math.Sqrt(3) / 2)
	v = ab * h / 3

	fmt.Printf("O VOLUME DA PIRÂMIDE É = %.2f METROS CÚBICOS\n", math.Round(v*100)/100)
}
