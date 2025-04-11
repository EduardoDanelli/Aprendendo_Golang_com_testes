package main

import "testing"

func TestArea(t *testing.T) {
	testesArea := []struct {
		nome     string
		forma    Forma
		esperado float64
	}{
		{forma: Retangulo{Largura: 12, Altura: 6}, esperado: 72.0},
		{forma: Circulo{Raio: 10}, esperado: 314.1592653589793},
		{forma: Triangulo{Base: 12, Altura: 6}, esperado: 36.0},
	}

	for _, teste := range testesArea {
		t.Run(teste.nome, func(t *testing.T) {
			resultado := teste.forma.Area()
			if resultado != teste.esperado {
				t.Errorf("Resultado: %#v, esperado: %v", resultado, teste.esperado)
			}
		})
	}
}
