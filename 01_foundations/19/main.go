package main

import (
	"fmt"
	"soma-go/matematica"

	"github.com/google/uuid"
)

func main() {
	s := matematica.Soma(10, 20)
	fmt.Println("O resultado é", s)
	fmt.Println("uuid", uuid.New().String())
}
