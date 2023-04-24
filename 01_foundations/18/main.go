package main

type myNumber int

type Number interface {
	~int | float32
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}

	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}

	return false
}

func main() {
	m := map[string]int{"claudio": 10000, "dalva": 6000, "antonio": 3000}
	m2 := map[string]float32{"claudio": 5300.30, "dalva": 4500.33, "antonio": 432.33}
	m3 := map[string]myNumber{"claudio": 5300, "dalva": 4500, "antonio": 432}
	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 10.0))
}
