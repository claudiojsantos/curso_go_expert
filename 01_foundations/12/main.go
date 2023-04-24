package main

import "fmt"

type Endereco struct {
	logradouro  string
	numero      string
	complemento string
	bairro      string
	cidade      string
	uf          string
}

type Cliente struct {
	nome    string
	idade   int
	address Endereco
	ativo   bool
}

func main() {
	cliente01 := Cliente{
		nome: "Claudio",
		address: Endereco{
			logradouro: "Rua Jose de Alencar",
			numero:     "436",
			bairro:     "Farol",
			cidade:     "Maceio",
			uf:         "AL",
		},
		idade: 48,
		ativo: true,
	}

	fmt.Println(cliente01.address)
}
