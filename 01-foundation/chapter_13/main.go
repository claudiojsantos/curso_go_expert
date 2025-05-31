package main

import "fmt"

type Client struct {
	Name   string
	Age    int
	Active bool
	Address
}

type Address struct {
	Street string
	Number int
	City   string
	State  string
	Zip    string
}

func (c Client) Deactivate() {
	c.Active = false
	fmt.Printf("Client %s deactivated\n", c.Name)
}

func main() {
	aldo := Client{
		Name:   "Aldo",
		Age:    5,
		Active: true,
		Address: Address{
			Street: "Rua do Aldinho",
			Number: 0,
			City:   "São Paulo",
			State:  "SP",
			Zip:    "00000-000",
		},
	}

	fmt.Println(aldo.Name)
	fmt.Println(aldo.Age)
	fmt.Println(aldo.Active)
	fmt.Println(aldo.Address.Street)
	fmt.Println(aldo.Address.Number)
	fmt.Println(aldo.Address.City)
	fmt.Println(aldo.Address.State)
	fmt.Println(aldo.Address.Zip)

	aldo.Deactivate()
}
