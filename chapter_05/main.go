package main

import "fmt"

func main() {
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	fmt.Printf("O valor do meu array é %d\n", meuArray[0])

	for i, v := range meuArray {
		fmt.Printf("O indice do meu array é %d e o valor é %d\n", i, v)
	}
}
