package main

import (
	"fmt"
)

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Println("Dentro do método salvar!")
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.nome)
}

func (u usuario) isMaiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Fulano", 20}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"Sicrano", 20}
	fmt.Println(usuario2)
	usuario2.salvar()
	fmt.Println(usuario2.isMaiorDeIdade())

	usuario2.fazerAniversario()
	fmt.Println(usuario2)
}
