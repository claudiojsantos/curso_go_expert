package main

import "fmt"

func main() {
	a := 100

	fmt.Println(a)
	fmt.Println(&a)

	var ponteiro *int = &a

	fmt.Println(ponteiro)
	fmt.Println(*ponteiro)
	*ponteiro = 150

	fmt.Println(*ponteiro)
	fmt.Println(a)
	fmt.Println(&a)

	a += 10

	fmt.Println(a)
	fmt.Println(*ponteiro)
}
