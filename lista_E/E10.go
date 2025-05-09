package main

import "fmt"

func main() {
	var N int
	var palindromo bool

	fmt.Scan(&N)

	vetor := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&vetor[i])
	}

	for i := 0; i < N; i++ {
		if vetor[i] == vetor[N-1-i] {
			palindromo = true
		} else {
			palindromo = false
			break
		}
	}

	if palindromo {
		fmt.Println("SIM")
	} else {
		fmt.Println("NAO")
	}
}
