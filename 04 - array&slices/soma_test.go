package main

import (
	"reflect"
	"testing"
)

func TestSoma(t *testing.T) {

	verificaMensagem := func(t *testing.T, resultado, esperado int, numeros []int) {
		t.Helper()
		if esperado != resultado {
			t.Errorf("esperado: '%d', resultado: '%d'", esperado, resultado)
		}
	}

	t.Run("arrays (possuem tamanho fixo)", func(t *testing.T) {

		numeros := [5]int{1, 2, 3, 4, 5}

		resultado := SomaArray(numeros)
		esperado := 15

		verificaMensagem(t, resultado, esperado, numeros[:])
	})

	t.Run("slices (não possuem tamanho fixo)", func(t *testing.T) {

		numeros := []int{1, 2, 3, 4, 5}

		resultado := SomaSlice([]int(numeros))
		esperado := 15

		verificaMensagem(t, resultado, esperado, numeros)

	})

}

func TestSomaTodoOResto(t *testing.T) {

	verificaSomas := func(t *testing.T, resultado, esperado []int) {
		t.Helper()
		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("resultado %v, esperado %v", resultado, esperado)
		}
	}

	t.Run("faz a soma de alguns slices", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{1, 2}, []int{0, 9})
		esperado := []int{2, 9}

		verificaSomas(t, resultado, esperado)
	})

	t.Run("soma slices vazios de forma segura", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{}, []int{3, 4, 5})
		esperado := []int{0, 9}

		verificaSomas(t, resultado, esperado)
	})
}
