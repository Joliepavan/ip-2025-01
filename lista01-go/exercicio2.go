package main

import (
	"fmt"
	"math"
)

func main() {
	var t, publico int
	var perc_p, perc_g, perc_a, perc_c float64
	var renda, p, g, a, c float64

	fmt.Scan(&t)
	for i := 1; i <= t; i++ {
		fmt.Scan(&publico, &perc_p, &perc_g, &perc_a, &perc_c)
		p = (perc_p / 100.0) * float64(publico)
		g = (perc_g / 100.0) * float64(publico)
		a = (perc_a / 100.0) * float64(publico)
		c = (perc_c / 100.0) * float64(publico)

		renda = p + (g * 5.00) + (a * 10.00) + (c * 20.00)
		fmt.Printf("A RENDA DO JOGO N. %d É = %.2f\n", i, math.Round(renda*100)/100)
	}
}
