package main

import (
	"Go.Course/OOP/clientes"
	"Go.Course/OOP/contas"
)

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	firstAccount := contas.ContaCorrente{Titular: clientes.Titutlar{Nome: "Test", CPF: "teste", Profissao: "Escrevente"}}

	secondAccount := contas.ContaCorrente{Titular: clientes.Titutlar{Nome: "Test2", CPF: "teste2", Profissao: "desenvolvedor"}}

	success, error := secondAccount.Transferir(100, &firstAccount)

	if success {
		println("saldo: ", secondAccount.ObterSaldo())
		println("sucesso ao transferir o valor para a conta")
	} else {
		println("erro ao transferir: ", error.Error())
	}
}
