package main

import "fmt"

func main() {
	var vetor [10]int
	var somaPares, qntdImpares int
	var pares []int
	var impares []int

	fmt.Println("Digite 10 números inteiros:")

	for i := 0; i < 10; i++ {
		fmt.Scan(&vetor[i])

		if vetor[i]%2 == 0 {
			pares = append(pares, vetor[i])
			somaPares += vetor[i]
		} else {
			impares = append(impares, vetor[i])
			qntdImpares++
		}
	}

	fmt.Println("Números pares: ", pares)
	fmt.Println("Soma dos números pares: ", somaPares)
	fmt.Println("Números ímpares: ", impares)
	fmt.Println("Quantidade de números ímpares:", qntdImpares)
}
