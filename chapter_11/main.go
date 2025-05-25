package main

import "fmt"

type Client struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	aldo := Client{
		Name:   "Aldo",
		Age:    5,
		Active: true,
	}

	fmt.Println(aldo.Name)
	fmt.Println(aldo.Age)
	fmt.Println(aldo.Active)

	aldo.Active = false
	fmt.Println(aldo.Active)
}
