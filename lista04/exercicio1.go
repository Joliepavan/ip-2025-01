package main

import "fmt"

func main() {
	var vetor []int 
	var X int       
	var result string

	maior := false 

	
	for i := 0; i < 10; i++ {
		fmt.Scan(&X)
		vetor = append(vetor, X) 

		
		if X > 50 {
			maior = true
			fmt.Println(i, X) 
		}
	}


	if !maior {
		result = "Não há números maiores que 50"
		fmt.Println(result)
	}
}


