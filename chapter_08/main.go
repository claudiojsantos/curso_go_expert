package main

import "fmt"

func main() {
	fmt.Println(sum(3, 5))
	name, err := namePerson("John")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Name:", name)
	}
	name, err = namePerson("")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Name:", name)
	}
}

func sum(a, b int) int {
	return a + b
}

func namePerson(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("nome deve ser preenchido")
	}
	return name, nil
}
