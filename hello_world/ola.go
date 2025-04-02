package main

import "fmt"

const OlaEmPortugues = "Ola, "
const OlaEmEspanhol = "Hola, "
const OlaEmFrances = "Bonjour, "

func Ola(nome string, idioma string) string {

	if nome == "" {
		nome = "Mundo"
	}

	return prefixoSaudacao(idioma) + nome
}

func prefixoSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case "frances":
		prefixo = OlaEmFrances
	case "espanhol":
		prefixo = OlaEmEspanhol
	default:
		prefixo = OlaEmPortugues
	}
	return
}

func main() {

	fmt.Println(Ola("Edu", ""))
}
