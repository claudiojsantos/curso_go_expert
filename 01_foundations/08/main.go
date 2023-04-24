package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, error := sum(2, 2)

	if error != nil {
		fmt.Println(error)
	}

	fmt.Printf("total %d\n", valor)
}

func sum(a, b int) (int, error) {
	if a+b < 50 {
		return 0, errors.New("O valor deve totalizar mais que 50")
	}

	return a + 5, nil
}
