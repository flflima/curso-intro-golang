package main

import "fmt"

func main() {
	canal := make(chan string, 2) // coloca um limite como o buffer

	canal <- "Olá Mundo!"

	mensagem := <-canal

	fmt.Println(mensagem)
}
