package main

import "fmt"

func sum(a, b *int) int {
	*a = 15
	*b = 25

	return *a + *b
}

func main() {
	var (
		valor01 = 10
		valor02 = 10
	)

	fmt.Println(sum(&valor01, &valor02))

	fmt.Println(valor01)
	fmt.Println(valor02)

}
