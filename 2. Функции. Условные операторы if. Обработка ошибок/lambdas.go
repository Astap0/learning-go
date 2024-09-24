package main

import (
    "fmt"
    "time"
)

// Функция, принимающая анонимную функцию в качестве аргумента
func processNumbers(a, b int, operation func(int, int) int) int {
    return operation(a, b)
}

// Возвращает анонимную функцию, которая замыкает переменную count
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    // Пример 1: Простая анонимная функция
    func() {
        fmt.Println("Привет из анонимной функции!")
    }()

    // Пример 2: Анонимная функция с параметрами
    sum := func(a, b int) int {
        return a + b
    }

    result := sum(3, 4)
    fmt.Printf("Сумма 3 и 4 равна: %d\n", result)

    // Пример 3: Анонимная функция как аргумент
    result = processNumbers(5, 10, func(x, y int) int {
        return x + y
    })
    fmt.Printf("Результат операции: %d\n", result)

    // Пример 4: Замыкания (Closures)
    increment := counter()

    fmt.Println(increment()) // 1
    fmt.Println(increment()) // 2
    fmt.Println(increment()) // 3

    // Пример 5: Анонимная горутина
    go func() {
        fmt.Println("Анонимная функция в горутине!")
    }()

    // Даем время горутине завершиться
    time.Sleep(1 * time.Second)
}
