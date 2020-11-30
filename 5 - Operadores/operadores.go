package main

import (
	"fmt"
)

func main() {
	// aritmeticos: + - * / %
	soma := 1 + 2
	subtracao := 2 - 3
	multiplicacao := 3 * 3
	divisao := 4 / 2
	resto := 10 % 5

	fmt.Println(soma, subtracao, multiplicacao, divisao, resto)

	// var numero1 int16 = 10
	// var numero2 int32 = 20
	// soma2 := numero2 + numero2  - erro pois são tipos diferentes 

	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	fmt.Println(true && false)
	fmt.Println(false || true)
	fmt.Println(!false)

	numero := 10
	numero++
	fmt.Println(numero)

	numero += 15
	fmt.Println(numero)

	// ++numero - não existe em go
	// não existe operador ternario em go
	
}