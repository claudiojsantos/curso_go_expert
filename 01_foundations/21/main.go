package main

func main() {
	a := 1
	b := 2
	c := 3

	if a > b {
		println(a)
	} else {
		println(b)
	}

	if a < b && c > 1 {
		println("a > b && c > 1")
	}

	a = 20

	switch a {
	case 1:
		println(1)
	case 2:
		println(2)
	default:
		println("nenhuma das opções")
	}
}
