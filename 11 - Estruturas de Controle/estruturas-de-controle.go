package main

import (
	"fmt"
)

func main() {
	fmt.Println("Estruturas de Controle")

	numero := -5

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 { // if init
		fmt.Println("Número é maior que zero")
	} else if numero < -10 {
		fmt.Println("Número é menor que -10")
	} else {
		fmt.Println("Entre zero e -10")
	}

	// fmt.Println(outroNumero) não pode ser acessada por aqui (escopo limitado ao IF/ELSE)
}
