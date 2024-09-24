package main

import "fmt"

func main() {
// 	printMessage("Вызов 1")
// 	printMessage("Вызов 2")
// 	printMessage(sayHello("Максим"))
	var age int = 70
	message, _ := enterTheClub(age)
	fmt.Printf("%s", message)
}

// func printMessage(message string) {
// 	fmt.Println(message)
// }

// func sayHello(name string) string /*Возвращаемое значение*/{

// 	return "Привет, "+ name
// }

func enterTheClub(age int) (string, bool) {
	if age >=18 && age < 45 {
		return "Входи, только акуратно", true
	} else if age >= 45 && age < 65 {
		return "Вам точно понравится эта музыка", true
	} else if age >= 65 {
		return "Это уже слишком для вас", false
	}
	return "Тебе нет 18-ти", false
}