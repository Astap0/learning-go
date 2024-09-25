package main

import "fmt"

// Обобщенная функция для суммирования чисел
func sum[T int | float64](numbers []T) T {
	var total T // Переменная для хранения суммы
	for _, number := range numbers {
		total += number // Добавляем каждый элемент к общей сумме
	}
	return total // Возвращаем сумму
}

func main() {
	// Пример с целыми числами
	intNumbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Сумма целых чисел:", sum(intNumbers)) // Выведет 15

	// Пример с числами с плавающей точкой
	floatNumbers := []float64{1.1, 2.2, 3.3}
	fmt.Println("Сумма дробных чисел:", sum(floatNumbers)) // Выведет 6.6
}