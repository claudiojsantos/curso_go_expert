package main

const a = "constante a"

// Variáveis Globais
var (
	b string  = "Variavel global b"
	c int     = 10
	d bool    = true
	e float32 = 10.5
)

func main() {
	var f string = "Variavel local f"

	println(a)
	println(b)
	println(c)
	println(d)
	println(e)
	println(f)
}
