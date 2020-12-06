package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("Olá Mundo!")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("Passando valor")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido --> %s", texto)
	}("Passando valor")
	fmt.Println(retorno)
}
