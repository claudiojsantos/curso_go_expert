package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 200}

	res, err := json.Marshal(conta)

	if err != nil {
		println("Erro: ", err)
	}

	println(string(res))

	// encode

	err = json.NewEncoder(os.Stdout).Encode(conta)

	if err != nil {
		println(err)
	}

	// inserir no struct

	jsonPuro := []byte(`{"n": 2, "s": 300}`)

	var contaX Conta

	err = json.Unmarshal(jsonPuro, &contaX)

	if err != nil {
		println(err)
	}

	println(contaX.Saldo)

}
