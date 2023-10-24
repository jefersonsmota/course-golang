package main

import (
	"course-golang/03-agencia-bancaria/clientes"
	"course-golang/03-agencia-bancaria/contas"
	"course-golang/03-agencia-bancaria/interfaces"
	"fmt"
)

func main() {
	cliente1 := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "John",
			CPF:       "12345678900",
			Profissao: "Marceneiro",
		},
		NumeroAgencia: 0001,
		NumeroConta:   12345,
	}
	cliente1.Depositar(1000.00)

	fmt.Printf("Cliente: %s, Saldo: R$ %.2f \n", cliente1.Titular.Nome, cliente1.ObterSaldo())
	//fmt.Println(cliente1)
	interfaces.PagarBoleto(&cliente1, 500)
	fmt.Printf("Cliente: %s, Saldo: R$ %.2f \n", cliente1.Titular.Nome, cliente1.ObterSaldo())

	cliente2 := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Marie",
			CPF:       "12345678901",
			Profissao: "Empresaria",
		},
		NumeroAgencia: 0001,
		NumeroConta:   12345,
	}
	cliente2.Depositar(1000.00)

	fmt.Printf("Cliente: %s, Saldo: R$ %.2f \n", cliente2.Titular.Nome, cliente2.ObterSaldo())

	//fmt.Println(cliente2)
}
