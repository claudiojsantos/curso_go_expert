package main

import "fmt"

func main() {
	var minha_var interface{} = "minha var\n"

	fmt.Printf(minha_var.(string))

	res, ok := minha_var.(int)

	fmt.Printf("O valor de res é %v e o valor de ok é %v", res, ok)
}
