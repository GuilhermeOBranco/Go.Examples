package main

import (
	"Go.Course/OOP/clientes"
	"Go.Course/OOP/contas"
)

func main() {
	contaSilvia := contas.ContaCorrente{Titular: clientes.Titutlar{Nome: "Geovanna", CPF: "teste", Profissao: "Escrevente"}}

	contaGui := contas.ContaCorrente{Titular: clientes.Titutlar{Nome: "Guilherme", CPF: "teste2", Profissao: "desenvolvedor"}}

	success, error := contaGui.Transferir(100, &contaSilvia)

	if success {
		println("saldo: ", contaGui.ObterSaldo())
		println("sucesso ao transferir o valor para a conta")
	} else {
		println("erro ao transferir: ", error.Error())
	}
}
