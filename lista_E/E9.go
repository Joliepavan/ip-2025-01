package main

import "fmt"

func main() {
	var N int

	fmt.Scan(&N)

	vetor := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&vetor[i])
	}

	soma := 0
	for i := 0; i < N; i++ {
		soma += vetor[i]
		if i == N-1 {
			fmt.Printf("%d\n", soma)
		} else {
			fmt.Printf("%d ", soma)
		}
	}
}
