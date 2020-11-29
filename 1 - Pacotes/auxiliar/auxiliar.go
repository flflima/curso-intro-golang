package auxiliar

import (
	"fmt"
)

// Escrever Colocando um comentario pois é boa prática em funções que são exportafas
func Escrever() { // Em maiuscula para ser "publica"
	fmt.Println("Escrevendo do pacote auxiliar")
	escrever2() // visivel pq estão no mesmo pacote
}