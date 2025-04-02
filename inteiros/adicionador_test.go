package main

import "testing"

func TestAdicionador(t *testing.T) {

	soma := Adicionador(2, 2)
	esperado := 4

	if soma != esperado {
		t.Errorf("resultado: '%d', esperado: '%d'", soma, esperado)
	}
}
