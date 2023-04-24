package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2323, 122, 322, 111, 245))
}

func sum(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}
