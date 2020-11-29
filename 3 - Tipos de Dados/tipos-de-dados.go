package main

import (
	"errors"
	"fmt"
)

func main() {
	// int8, int16, int32, int64
	// suportam numeros negativos
	var numero int16 = 100
	fmt.Println(numero)

	// int (usa a arquitetura do pc como base
		//  se pc 32 bits, usa int32
		//  se pc 64 bits, usa bit64), 
	numeroInferencia := 10
	fmt.Println(numeroInferencia)

	// uint - unsigned int
	var numero2 uint32 = 100
	fmt.Println(numero2)

	// Alias
	// int32 == rune (usado principalmente com caracteres)
	var numero3 rune = 1234
	fmt.Println(numero3)

	// uint8 = byte
	var numero4 byte = 123
	fmt.Println(numero4)

	// float32, float64
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123323424.433
	fmt.Println(numeroReal2)

	// float: mesmo esquema do int em relação a arquitetura do pc
	numeroReal3 := 1234.44
	fmt.Println(numeroReal3)


	// não existe char

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char) // imprime o valor na tabela ASCII


	// valor zero: quando não se atribui nenhum valor a declaração abaixo
	var texto int16
	fmt.Println(texto)


	var booleano1 bool = true
	fmt.Println(booleano1)


	// valor zero de error é <nil>
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}