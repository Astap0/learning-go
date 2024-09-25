package main

import "fmt"

func main() {
	users := map[string]int{
		"Vasya":12,
		"Petya":23,
		"Kostya":48,
	}
for key, value := range users {
	fmt.Println(key,":",value)
}
}