package main

import "fmt"

func Adicionador(valor1, valor2 int) (soma int) {
	soma = valor1 + valor2
	return soma
}

func main() {

	fmt.Println(Adicionador(2, 2))
}
