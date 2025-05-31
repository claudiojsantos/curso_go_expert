package main

import "fmt"

func main() {
	number_a := 20
	number_b := 30
	fmt.Println(sum(&number_a, &number_b))
	fmt.Println(number_a)
	fmt.Println(number_b)
}

func sum(a, b *int) int {
	// a = 10 // resultado de number_a será 20
	*a = 10 // resultado de number_a será 10
	*b = 90 // resultado de number_b será 90
	return *a + *b
}
