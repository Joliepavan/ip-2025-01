package main

import (
	"fmt"
	"math"
)

func main() {
	var n1, n2, n3 float64
	var X, Y float64

	fmt.Println("Insira três números entre 1 e 9")
	fmt.Scan(&n1, &n2, &n3)

	X = n1*100 + n2*10 + n3
	Y = math.Pow(X, 2.0)

	if n1 < 1 || n1 > 9 || n2 > 9 || n3 > 9 {
		fmt.Println("DIGITO INVALIDO")
	} else {
		fmt.Printf("%.0f, %.0f\n", X, Y)
	}
}
