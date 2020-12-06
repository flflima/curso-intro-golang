package main

import (
	"fmt"
)

func closure() func() {
	texto := "Dentro da função closure"

	// funcao := func() {
	// 	fmt.Println(texto)
	// }

	// return funcao

	return func() {
		fmt.Println(texto)
	}
}

func main() {
	texto := "Dentro da função main"
	fmt.Println(texto)

	// funcaoNova := closure()
	// funcaoNova()
	closure()()
}
