package main

import (
	"fmt"
)

func main() {
	var N int
	var v1, v2 int
	var pe, result int
	vetor1 := make([]int, 0)
	vetor2 := make([]int, 0)

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&v1)
		vetor1 = append(vetor1, v1)
	}

	for i := 0; i < N; i++ {
		fmt.Scan(&v2)
		vetor2 = append(vetor2, v2)
	}

	for i := 0; i < N; i++ {
		pe = vetor1[i] * vetor2[i]
		result += pe
	}
	fmt.Print(result)
}
