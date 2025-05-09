package main

import "fmt"

func main() {

	var N int
	var vetor []int
	var valor int

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&valor)
		vetor = append(vetor, valor)
	}

	var maiorsequencia int = 1
	var sequencia int = 1

	for i := 1; i < N; i++ {
		if vetor[i] > vetor[i-1] {
			sequencia++
			if sequencia > maiorsequencia {
				maiorsequencia = sequencia
			}
		} else {
			sequencia = 1
		}
	}

	fmt.Println(maiorsequencia)
}
