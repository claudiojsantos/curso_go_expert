package main

import "fmt"

func main() {
	arraySlice := []int{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000}

	fmt.Printf("Tamanho %d, Capacidade %d, Array %v\n", len(arraySlice), cap(arraySlice[:0]), arraySlice)
	fmt.Printf("Tamanho %d, Capacidade %d, Array %v\n", len(arraySlice[:2]), cap(arraySlice[:2]), arraySlice[:2])
	fmt.Printf("Tamanho %d, Capacidade %d, Array %v\n", len(arraySlice[2:]), cap(arraySlice[2:]), arraySlice[2:])

	arraySlice = append(arraySlice, 1100)

	fmt.Printf("Tamanho %d, Capacidade %d, Array %v\n", len(arraySlice), cap(arraySlice[:0]), arraySlice)
}
