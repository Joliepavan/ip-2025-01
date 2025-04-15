package main

import "fmt"

func main() {
	var n, Aprovados, Reprovados, Exame int
	var nota1, nota2 float64

	fmt.Scan(&n)

	for i := 1; i <= n; i++ {

		fmt.Scan(&nota1, &nota2)

		media := (nota1 + nota2) / 2

		if media < 3 {
			fmt.Printf("Aluno %d :Reprovado \n", i)
			Reprovados++
		} else if media >= 3 && media <= 7 {
			fmt.Printf("Aluno %d : Exame \n", i)
			Exame++
		} else if media > 7 && media <= 10 {
			fmt.Printf("Aluno %d : Aprovado \n", i)
			Aprovados++
		}

	}
	fmt.Printf("Total Aprovados: %d \n", Aprovados)
	fmt.Printf("Total Exame: %d \n", Exame)
	fmt.Printf("Total Reprovados: %d \n", Reprovados)

}
