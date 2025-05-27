package main

import "fmt"

var y interface{} = "Hi Go"

func main() {
	fmt.Printf(y.(string))
}
