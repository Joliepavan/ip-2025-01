package main

import (
	"fmt"
	"sort"
)

func main() {

	var N int
	var vetor []int
	var valor int

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&valor)
		vetor = append(vetor, valor)
	}
	sort.Ints(vetor)
	if N > 1 {
		for i := 0; i < N; i++ {
			fmt.Print(vetor[i])
			if i != N-1 {
				fmt.Print(" ")
			}
		}
	} else {
		fmt.Print(vetor[0])
	}

}
