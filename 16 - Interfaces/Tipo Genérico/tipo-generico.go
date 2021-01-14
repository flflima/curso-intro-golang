package main

import (
	"fmt"
)

func generica(interf interface{}) { // qualquer tipo atende esta interface
	fmt.Println(interf)
}

func main() {
	generica("string")
	generica(100)
	generica(false)
}
