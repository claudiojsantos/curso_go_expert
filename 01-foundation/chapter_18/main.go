package main

import "fmt"

var x interface{} = 20
var y interface{} = "Hi Go"

func main() {
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)
}
