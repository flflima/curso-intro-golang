package main

import (
	"reflect"
	"fmt"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int
	array1[0] = 100
	fmt.Println(array1)

	array2 := [5]string{"-1-", "-2-", "-3-", "-4-", "-5-"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	//==================

	slice := []int{1, 2, 3}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 4)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)
}
