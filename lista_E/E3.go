package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	var valores int
	var vetor []int

	for i := 0; i < N; i++ {
		fmt.Scan(&valores)
		vetor = append(vetor, valores)
	}

	for i := 0; i < N; i++ {
		if N == 1 {
			fmt.Print("0")
		} else if i == 0 {
			fmt.Print(vetor[1])
		} else if i == N-1 {
			fmt.Print(vetor[N-2])
		} else {
			fmt.Print(vetor[i-1] + vetor[i+1])
		}

		if i != N-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
