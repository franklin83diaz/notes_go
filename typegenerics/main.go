package main

import (
	"fmt"
	"reflect"
)

func printt[T any](a T, b T) {
	fmt.Println(a, b)
	return
}

func check[T any](x T) {
	// Usamos reflexión para obtener el tipo de 'a'
	if reflect.TypeOf(x).Kind() == reflect.Int {
		fmt.Println("Is init val")
	} else {
		fmt.Println("Is not init val")
	}

	return
}

func main() {
	fmt.Println("example type generics")
	printt(1, 2)
	printt("hello", "world")
	fmt.Println()
	check(1)
	check[string]("hello")
}
