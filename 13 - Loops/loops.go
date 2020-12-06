package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Loops")

	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i")
	// 	time.Sleep(time.Second)
	// }

	// fmt.Println(i)

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando j")
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"João", "Davi", "Pedro"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for _, letra := range "PALAVRA" {
		fmt.Println(string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
