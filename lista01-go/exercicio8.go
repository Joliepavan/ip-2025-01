package main

import (
	"fmt"
	"math"
)

func main() {
	var r, h, at, ac, al, x float64
    fmt.Print("insira o raio da base e a altura do cilindro \n")
	fmt.Scan(&r)
	fmt.Scan(&h)

	ac = math.Pow(r, 2.0) * 3.14159
	al = 2 * r * h * 3.14159
	at = 2*ac + al
	x = at * 100

	fmt.Printf("O VALOR DO CUSTO É= %.2f\n", math.Round(x*100)/100)
}
