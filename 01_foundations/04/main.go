package main

import "fmt"

type ID int

var (
	f ID = 1
)

func main() {
	fmt.Printf("O tipo de f é %T\n", f)
	fmt.Printf("O valor de f é %v", f)
}
