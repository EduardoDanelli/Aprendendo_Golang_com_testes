package main

import "testing"

func TestOla(t *testing.T) {

	verificaMensagem := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado: '%s', esperado: '%s'", resultado, esperado)
		}
	}

	t.Run("diz ola para as pessoas", func(t *testing.T) {
		resultado := Ola("Edu")
		esperado := "Ola, Edu"

		verificaMensagem(t, resultado, esperado)
	})

	t.Run("diz 'Ola, mundo' quando uma string vazia for passada", func(t *testing.T) {
		resultado := Ola("")
		esperado := "Ola, Mundo"

		verificaMensagem(t, resultado, esperado)
	})
}
