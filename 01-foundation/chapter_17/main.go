package main

import "fmt"

type Conta struct {
	agency  int
	account int
	balance float64
}

func (c *Conta) deposit(value float64) float64 {
	c.balance += value
	return c.balance
}

func (c *Conta) withdraw(value float64) float64 {
	c.balance -= value
	return c.balance
}

func main() {
	conta := Conta{
		agency:  1,
		account: 1,
		balance: 1000.00,
	}

	fmt.Printf("Saldo: %.2f\n", conta.balance)
	conta.deposit(150)
	fmt.Printf("Saldo Atualizado: %.2f\n", conta.balance)

	conta.withdraw(200)
	fmt.Printf("Saldo Atualizado 2: %.2f\n", conta.balance)
}
