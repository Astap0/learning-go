package main

import (
	"fmt"
)

// 1. Простые типы
var intVar int = 10
var floatVar float64 = 3.14
var boolVar bool = true
var stringVar string = "Hello, Go!"

// 2. Составные типы
var arrayVar [5]int = [5]int{1, 2, 3, 4, 5}
var sliceVar []int = []int{1, 2, 3, 4, 5}

type Person struct {
	Name string
	Age  int
}

var personVar = Person{Name: "Alice", Age: 30}

var mapVar = map[string]int{
	"one": 1,
	"two": 2,
	"three": 3,
}

// 3. Функциональные типы
type MyFunc func(int) int

func square(x int) int {
	return x * x
}

// 4. Интерфейсы
type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

// 5. Указатели
func increment(val *int) {
	*val++
}

func main() {
	// Вывод простых типов
	fmt.Println("Простые типы:")
	fmt.Println("intVar:", intVar)
	fmt.Println("floatVar:", floatVar)
	fmt.Println("boolVar:", boolVar)
	fmt.Println("stringVar:", stringVar)

	// Вывод составных типов
	fmt.Println("\nСоставные типы:")
	fmt.Println("arrayVar:", arrayVar)
	fmt.Println("sliceVar:", sliceVar)
	fmt.Println("personVar:", personVar)
	fmt.Println("mapVar:", mapVar)

	// Использование функциональных типов
	fmt.Println("\nФункциональные типы:")
	var myFunc MyFunc = square
	fmt.Println("Square of 4:", myFunc(4))

	// Использование интерфейсов
	fmt.Println("\nИнтерфейсы:")
	var animal Animal = Dog{}
	fmt.Println("Dog speaks:", animal.Speak())

	// Использование указателей
	fmt.Println("\nУказатели:")
	count := 5
	fmt.Println("Before increment:", count)
	increment(&count)
	fmt.Println("After increment:", count)
}
