package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Jeferson"
	idade := 0
	versao := 1.1
	println("Olá Sr.", nome)
	println("Versão ", versao)
	println("tipo da varivael nome", reflect.TypeOf(nome).String())

	fmt.Printf("Informe idade: ")
	fmt.Scan(&idade)

	if idade <= 18 {
		fmt.Println("menor")
	} else if idade >= 18 {
		fmt.Println("maior")
	}

	pontosPlanningPoker := []int{1, 2, 3, 5, 8, 13, 21}
	fmt.Println("pontosPlanningPoker:", pontosPlanningPoker)
	fmt.Println(cap(pontosPlanningPoker))
	pontosPlanningPoker = append(pontosPlanningPoker, 40)
	fmt.Println("pontosPlanningPoker:", pontosPlanningPoker)
	fmt.Println(cap(pontosPlanningPoker))
	pontosPlanningPoker = append(pontosPlanningPoker, 52)
	fmt.Println("pontosPlanningPoker:", pontosPlanningPoker)
	fmt.Println(cap(pontosPlanningPoker))
}
