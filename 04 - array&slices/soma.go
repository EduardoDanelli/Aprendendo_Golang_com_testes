package main

func SomaArray(numeros [5]int) int {
	var soma int
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

func SomaSlice(numeros []int) int {
	var soma int
	for _, numero := range numeros {
		soma += numero
	}

	return soma
}

func SomaTodoOResto(numerosParaSomar ...[]int) []int {
	var somas []int
	for _, numeros := range numerosParaSomar {
		if len(numeros) == 0 {
			somas = append(somas, 0)
		} else {
			final := numeros[1:]
			somas = append(somas, SomaSlice(final))
		}
	}

	return somas
}
