package main

import "fmt"

var (
	meuArray [3]int
)

func main() {
	meuArray[0] = 100
	meuArray[1] = 200
	meuArray[2] = 300

	for i, v := range meuArray {
		fmt.Printf("O indice é %d e seu valor é %d\n", i, v)
	}
}
