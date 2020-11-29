package main

import (
	"fmt"
)

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println("Função f")
		return txt
	}
	
	retorno := f("Texto da função")
	fmt.Println(retorno)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15) // dois retornos então atribui a duas variaveis
	fmt.Println(resultadoSoma, resultadoSubtracao)

	resultadoSomaNovo, _ := calculosMatematicos(20, 10) // desconsidera a segunda variavel para não gerar o erro
	fmt.Println(resultadoSomaNovo)
}