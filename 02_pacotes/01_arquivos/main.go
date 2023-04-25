package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")

	if err != nil {
		panic(err)
	}

	//tamanho, err := f.WriteString("Hello, World!!!")
	tamanho, err := f.Write([]byte("Hello, World!!!"))

	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso com o tamamnho de %d bytes\n", tamanho)

	f.Close()

	arquivo, err := os.ReadFile("arquivo.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(arquivo) // Byte
	fmt.Println(string(arquivo))

	// Caso o tamanho do arquivo seja grande ou muitas vezes maior que a memória

	arquivo2, err := os.Open("arquivo.txt")

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)

		if err != nil {
			break
		}

		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("arquivo.txt")

	if err != nil {
		panic(err)
	}
}
