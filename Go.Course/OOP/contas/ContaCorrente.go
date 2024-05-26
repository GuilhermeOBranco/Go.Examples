package contas

import (
	"errors"

	"Go.Course/OOP/clientes"
)

type ContaCorrente struct {
	Titular       clientes.Titutlar
	NumeroAgencia string
	NumeroConta   string
	saldo         float64
}

func (conta *ContaCorrente) Sacar(valorSaque float64) string {

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

func (conta *ContaCorrente) Depositar(valorDeposito float64) (saldoConta float64, err error) {

	if valorDeposito < 0 {
		return 0, errors.New("apenas valores positivos podem ser depositados")
	}

	conta.saldo += valorDeposito

	return conta.saldo, nil
}

func (contaOrigem *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) (success bool, err error) {

	if valorTransferencia < 0 {
		return false, errors.New("Nao pode ser realizado uma transferencia de valores negativos")
	}

	if contaOrigem.saldo < valorTransferencia {
		return false, errors.New("Nao existe saldo o suficiente para a transferencia")
	}

	contaOrigem.saldo -= valorTransferencia

	_, depoError := contaDestino.Depositar(valorTransferencia)

	if depoError != nil {
		//retorna o dinheiro para a conta de origem
		contaOrigem.saldo += valorTransferencia

		return false, err
	}

	return true, nil

}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
