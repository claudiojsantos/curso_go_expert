package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []string{"um", "dois", "tres"}

	for k, v := range numeros {
		println(k, v)
	}

	for _, v := range numeros {
		println(v)
	}

	for k, _ := range numeros {
		println(k)
	}

	i := 0

	for i < 10 {
		println(i)
		i++
	}

	for {
		println("Hello World")
	}
}
