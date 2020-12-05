package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ponteiros")

	var variavel int
	var ponteiro *int

	fmt.Println(variavel, ponteiro)

	variavel = 100
	ponteiro = &variavel

	fmt.Println(variavel, ponteiro)
	fmt.Println(variavel, *ponteiro)

	variavel++ // incrementa as duas variaveis, pois é um ponteiro

	fmt.Println(variavel, *ponteiro)
}
