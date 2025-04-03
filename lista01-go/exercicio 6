package main

import (
	"fmt"
	"math"
)

func main() {
	var n, i int
	var x, y float64
	i = 1
	fmt.Print("Digite a quantidade de temperaturas a serem convertidas: ")
	fmt.Scan(&n)

	for i <= n {
		fmt.Print("Digite a temperatura em Fahrenheit: ")
		fmt.Scan(&x)
		y = 5 * (x - 32) / 9
		fmt.Printf("%.2f FAHRENHEIT EQUIVALE A %.2f CELSIUS\n", math.Round(x*100)/100, math.Round(y*100)/100)
		i++
	}
}
