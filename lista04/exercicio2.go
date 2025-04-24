package main

import "fmt"

func main(){
	var v1 [10]int
	var v2 [5]int
	var pares []int
	var impares []int

	fmt.Println("Digite o primeiro vetor: ")
	for i := 0; i < 10; i++ {
		fmt.Scan(&v1[i])
	}

	fmt.Println("Digite o segundo vetor: ")
	for i := 0; i < 5; i++ {
		fmt.Scan(&v2[i])
	}

	somaV2 := 0
	for i := 0; i < 5; i++ {
		somaV2 += v2[i]
	}

	for i := 0; i < 10; i++ {
		if v1[i]%2 == 0 {
			pares = append(pares, v1[i]+somaV2)
		} else {
			impares = append(impares, v1[i]+somaV2)
		}
	}
  
	fmt.Println("Primeiro vetor resultante: ", pares)
	fmt.Println("Segundo vetor resultante :", impares)
}
