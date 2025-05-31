package main

import "fmt"

func main() {
	salarios := map[string]float32{
		"John": 1000,
		"Jane": 2000,
		"Jim":  3000,
	}

	salarios_ti := make(map[string]float32)
	salarios_ti["John"] = 5000
	salarios_ti["Jane"] = 6000
	salarios_ti["Jim"] = 7000

	for name, salario := range salarios {
		fmt.Printf("O salário de %s é %.2f\n", name, salario)
	}

	for _, salario := range salarios_ti {
		fmt.Printf("Salário %.2f\n", salario)
	}
}
