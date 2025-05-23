package main

import "fmt"

type ID int

var (
	i ID = 10
)

func main() {
	fmt.Printf("O tipo de i é %T\n", i)
	fmt.Printf("O valor de i é %v\n", i)
}
