package main

import "fmt"

type Cliente struct {
	nome  string
	idade int
	ativo bool
}

func main() {
	cliente01 := Cliente{
		nome:  "Claudio",
		idade: 48,
		ativo: true,
	}

	fmt.Println(cliente01.nome)
}
