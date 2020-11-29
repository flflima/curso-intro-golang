package main

import (
	"fmt"
)

func main() {
	var variavel1 string = "Variável 1" // declaração explicita
	variavel2 := "Variável 2" // declaração implicita (inferencia de tipo)
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "lala"
		variavel4 string = "lele"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "ops", "foi mal"
	fmt.Println(variavel5, variavel6)

	// inversão de valores
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)

	variavel1 = "mudei"
	fmt.Println(variavel1)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)
	// constante1 = "novo valor" --> erro 
}