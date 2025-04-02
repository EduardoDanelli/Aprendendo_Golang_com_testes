package main

import "fmt"

const OlaEmPortugues = "Ola, "

func Ola(nome string) string {

	if nome == "" {
		nome = "Mundo"
	}

	return OlaEmPortugues + nome
}

func main() {

	fmt.Println(Ola("Edu"))
}
