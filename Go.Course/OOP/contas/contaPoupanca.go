package contas

import (
	"errors"

	"Go.Course/OOP/clientes"
)

type ContaPoupanca struct {
	Titular       clientes.Titutlar
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (conta *ContaPoupanca) Sacar(valorSaque float64) string {

	if valorSaque < 0 {
		return "apenas valores positivos podem ser sacados"
	}

	podeSacar := valorSaque < conta.saldo

	if podeSacar {
		conta.saldo -= valorSaque
		return "saque realizado com sucesso"
	}

	return "nao foi possivel realizar o saque"

}

func (conta *ContaPoupanca) Depositar(valorDeposito float64) (saldoConta float64, err error) {

	if valorDeposito < 0 {
		return 0, errors.New("apenas valores positivos podem ser depositados")
	}

	conta.saldo += valorDeposito

	return conta.saldo, nil
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
