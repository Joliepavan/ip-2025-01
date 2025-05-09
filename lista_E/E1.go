package main

import "fmt"

func main() {
	var atletas int
	var vetor []float64
	var alturas float64
	var soma float64

	fmt.Scan(&atletas)

	for i := 0; i < atletas; i++ {
		fmt.Scan(&alturas)
		vetor = append(vetor, alturas)

		soma = soma + alturas
	}

	media := soma / float64(atletas)

	for i := 0; i < atletas; i++ {
		if vetor[i] > media {
			fmt.Printf("%.2f\n", vetor[i])
		}
	}
}
