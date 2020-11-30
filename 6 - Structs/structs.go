package main

import (
	"fmt"
)

type usuario struct {
	nome     string
	idade    int8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "fulano"
	u.idade = 20
	fmt.Println(u)

	endereco := endereco{"Rua dos Bobos", 0}
	u2 := usuario{"beltrano", 21, endereco}
	fmt.Println(u2)

	u3 := usuario{nome: "sicrano"}
	fmt.Println(u3)
}
