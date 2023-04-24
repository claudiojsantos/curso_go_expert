package main

import "fmt"

func main() {
	total := func() int {
		return sum(1, 53, 2, 232, 66, 342, 22, 44)
	}()

	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}
