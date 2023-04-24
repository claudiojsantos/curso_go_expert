package main

import "fmt"

func main() {
	salarios := map[string]int{"paulo": 2000, "joao": 3000, "maria": 5000}

	for nome, salario := range salarios {
		fmt.Printf("%s = %d\n", nome, salario)
	}
}
