package main

func Repetir(caractere string, qntRepeticoes int) string {

	var repeticoes string
	for i := 0; i < qntRepeticoes; i++ {
		repeticoes += caractere
	}
	return repeticoes
}
