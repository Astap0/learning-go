package main

import (
	"fmt"
	"reflect"
)

func main() {
	message := "Я скоро стану GoLang Ninja"
	a, b, c := 1, 2, 3
	a, c, b = c, a, b
	// var message string = "Я скоро стану GoLang Ninja"
	// const message string = "Я скоро стану GoLang Ninja"
	/*
	Коммент
	*/
	fmt.Println(message)
	fmt.Println(reflect.TypeOf(message))
}