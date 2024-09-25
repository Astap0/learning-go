package main

import "fmt"

// Функция для работы с массивом
func printArray(arr [5]int) {
    fmt.Println("Массив внутри функции:")
    for i, v := range arr {
        fmt.Printf("Индекс %d: Значение %d\n", i, v)
    }
}

// Функция для работы с срезом
func printSlice(slice []int) {
    fmt.Println("Срез внутри функции:")
	slice[2] = 100
    for i, v := range slice {
        fmt.Printf("Индекс %d: Значение %d\n", i, v)
	slice = append(slice, 2,3)
    }
}

func main() {
    // 1. Работа с массивами
    var arr [5]int // Объявление массива на 5 элементов
    arr[0] = 10    // Инициализация элементов
    arr[1] = 20
    arr[2] = 30
    arr[3] = 40
    arr[4] = 50

    fmt.Println("Массив:", arr)

    // Передаем массив в функцию
    printArray(arr) // Массив остается неизменным

    // 2. Работа с срезами
    slice := []int{1, 2, 3} // Инициализация среза
    fmt.Println("Срез:", slice)

    // Добавляем элементы в срез
    slice = append(slice, 4, 5)
    fmt.Println("Срез после добавления элементов:", slice)

    // Передаем срез в функцию
    printSlice(slice) // Срез может изменяться внутри функции

    // Изменяем элемент среза
    slice[0] = 100
    fmt.Println("Срез после изменения:", slice)
}

/*
| Характеристика     | Массив                         | Срез                            |
|--------------------|--------------------------------|---------------------------------|
| Размер             | Фиксированный                  | Динамический                    |
| Передача           | Копируется целиком             | Передается по ссылке (указатель)|
| Инициализация      | `var arr [5]int`               | `slice := []int{}`              |
| Функции            | Не поддерживает `append()`     | Поддерживает `append()`         |
| Производительность | Может занимать больше памяти   | Более эффективен по памяти      |

### Итог
- **Массивы** подходят для фиксированных коллекций данных, но менее гибкие.
- **Срезы** более удобны и мощны для работы с динамическими коллекциями данных.
*/