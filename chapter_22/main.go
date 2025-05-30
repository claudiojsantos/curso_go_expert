package main

import "fmt"

func main() {
	nomes := []string{"João", "Maria", "José", "Pedro", "Ana"}

	for i, nome := range nomes {
		fmt.Println(i, nome)
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}
}
