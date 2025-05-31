package main

func main() {
	a := 30
	println(&a) // 0xc000046738
	b := a
	b = 40
	println(b) // 40
	println(a) // 30
	c := &a
	println(c) // 0xc000046738
	*c = 50
	println(*c) // 50
	println(a)  // 50
}
