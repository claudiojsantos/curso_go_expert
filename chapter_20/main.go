package main

import "fmt"

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var total T
	for _, v := range m {
		total += v
	}
	return total
}

func comparar[T comparable](a, b T) bool {
	return a == b
}

func main() {
	salarios := map[string]int{
		"John": 1000,
		"Jane": 2000,
		"Jim":  3000,
	}

	salarios2 := map[string]float64{
		"John": 10000.0,
		"Jane": 20000.50,
		"Jim":  30000.0,
	}

	fmt.Println(Soma(salarios))
	fmt.Println(Soma(salarios2))
	fmt.Println(comparar(1, 1))
}
