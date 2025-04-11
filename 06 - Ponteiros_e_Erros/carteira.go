package main

import (
	"errors"
	"fmt"
)

var ErroSaldoInsuficiente = errors.New("não é possível retirar: saldo insuficiente")

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Carteira struct {
	saldo Bitcoin
}

func (c *Carteira) Depositar(quantidade Bitcoin) {
	fmt.Printf("O endereço do saldo no Depositar é: %v\n", &c.saldo)

	c.saldo += quantidade
}

func (c *Carteira) Retirar(quantidade Bitcoin) error {
	fmt.Printf("O endereço do saldo no Retirar é: %v\n", &c.saldo)

	if c.saldo < quantidade {
		return ErroSaldoInsuficiente
	}

	c.saldo -= quantidade
	return nil
}

func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}
