package main

import (
	"fmt"
)

func main() {
	var i, n int
	i = 2
	fmt.Print("Digite um número: ")
	fmt.Scan(&n)

	if 5 < n && n < 2000 {
		for i <= n {
			fmt.Printf("%d ^ %d = %d\n", i, i, i*i)
			i += 2
		}
	}
}
