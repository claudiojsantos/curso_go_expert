package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Printf("O tamanho do slice é %d e sua capacidade é %d\n", len(s), cap(s))
	fmt.Printf("O slice %v tem tamanho de %d e capacidade de %d\n", s[:2], len(s[:2]), cap(s[:2]))
	fmt.Printf("O slice %v tem tamanho de %d e capacidade de %d\n", s[2:], len(s[2:]), cap(s[2:]))
	s = append(s, 110)
	fmt.Printf("O slice %v tem tamanho de %d e capacidade de %d\n", s, len(s), cap(s))
}
