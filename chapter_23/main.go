package main

import "fmt"

func main() {
	a := 11
	b := 2
	c := 10

	if a > b && a > c {
		fmt.Println("a é o maior")
	} else if b > a && b > c {
		fmt.Println("b é o maior")
	} else {
		fmt.Println("c é o maior")
	}

	switch a {
	case 3:
		fmt.Println("a é 3")
	case 4:
		fmt.Println("a é 4")
	default:
		fmt.Println("a não é 3 nem 4")
	}

}
