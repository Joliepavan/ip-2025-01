package main

import (
	"fmt"
)

func main() {
	var a, an, r, n, soma int

	fmt.Print("Digite os valores:")
	fmt.Scan(&a, &r, &n)

	an = a + (n-1)*r
	soma = n * (a + an) / 2

	fmt.Println(soma)
}
