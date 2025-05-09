package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	var vetor []float64
	var valores float64

	for i := 0; i < N; i++ {
		fmt.Scan(&valores)
		vetor = append(vetor, valores)

	}

	maximo := vetor[0]
	for i := 1; i < N; i++ {
		if vetor[i] > maximo {
			maximo = vetor[i]
		}
	}

	minimo := vetor[0]
	for i := 1; i < N; i++ {
		if vetor[i] < minimo {
			minimo = vetor[i]
		}
	}

	for i := 0; i < N; i++ {
		var normalizacao float64
		normalizacao = (vetor[i] - minimo) / (maximo - vetor[0])

		if vetor[i]-minimo == 0 || maximo-minimo == 0 {
			fmt.Print("0.00 ")
		} else {
			fmt.Printf("%.2f ", normalizacao)
		}
	}
}
