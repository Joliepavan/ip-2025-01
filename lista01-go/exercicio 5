package main

import (
	"fmt"
)

func main() {
	var u int
	var f, n float64
	var cons string

	fmt.Print("Digite o número da conta, consumo em kWh e tipo de consumidor (R, C ou I): ")
	fmt.Scan(&u, &f, &cons)

	fmt.Printf("CONTA: %d\n", u)

	if cons == "R" {
		n = 5.00 + (0.05 * f)
	} else if cons == "C" {
		if f <= 80 {
			n = 500.00
		} else {
			n = 500.00 + (0.25 * (f - 80))
		}
	} else if cons == "I" {
		if f<=100{
		    n=800.00
		}else {
		    n=800+ (0.04 * (f-100))
		}
