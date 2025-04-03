package main

import (
	"fmt"
	"math"
)

func main() {
	var salarioMinimo, qkw, custoKw, custoConsumo, custoDesc float64

	fmt.Print("Digite o salário mínimo: ")
	fmt.Scan(&salarioMinimo)
	fmt.Print("Digite a quantidade de kW consumidos: ")
	fmt.Scan(&qkw)

	custoKw = (salarioMinimo * 7) / 1000
	custoConsumo = custoKw * qkw
	custoDesc = custoConsumo - (custoConsumo * 0.1)

	fmt.Printf("Custo por kW: R$ %.2f\n", math.Round(custoKw*100)/100)
	fmt.Printf("Custo do consumo: R$ %.2f\n", math.Round(custoConsumo*100)/100)
	fmt.Printf("Custo com desconto: R$ %.2f\n", math.Round(custoDesc*100)/100)
}
