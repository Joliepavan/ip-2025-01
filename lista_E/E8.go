package main

import (
	"fmt"
)

func main() {
	var N int
	var numeros int
	count := 0
	vetor := make([]int, 0)

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&numeros)
		vetor = append(vetor, numeros)
	}

	for i := 0; i < N; i++ {
		unico := true
		for j := 0; j < i; j++ {
			if vetor[i] == vetor[j] {
				unico = false
				break
			}
		}
		if unico {
			count++
		}
	}
	fmt.Println(count)
}
