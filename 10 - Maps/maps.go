package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "fulano",
			"ultimo":   "de tals",
		},
		"curso": {
			"nome": "engenharia",
		},
	}
	fmt.Println(usuario2)

	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"tipo": "gemeos",
	}
	fmt.Println(usuario2)
}
