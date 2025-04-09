package main

import (
	"fmt"
)

func main() {
	var h, m, s, x1, x2, total int

	fmt.Print("Digite as horas, minutos e segundos: ")
	fmt.Scan(&h, &m, &s)

	x1 = h * 3600
	x2 = m * 60
	total = x1 + x2 + s

	fmt.Println("O TEMPO EM SEGUNDOS É:", total)
}
