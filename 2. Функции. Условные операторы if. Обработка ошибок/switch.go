package main

import (
    "fmt"
    "time"
)

func main() {
    // Пример 1: Простое использование switch с целыми числами
    dayOfWeek := time.Now().Weekday()

    fmt.Println("Пример 1: День недели")
    switch dayOfWeek {
    case time.Monday:
        fmt.Println("Сегодня понедельник.")
    case time.Tuesday:
        fmt.Println("Сегодня вторник.")
    case time.Wednesday:
        fmt.Println("Сегодня среда.")
    case time.Thursday:
        fmt.Println("Сегодня четверг.")
    case time.Friday:
        fmt.Println("Сегодня пятница.")
    case time.Saturday, time.Sunday: // несколько значений в одном case
        fmt.Println("Сегодня выходной!")
    default:
        fmt.Println("Неизвестный день недели.")
    }

    // Пример 2: Использование switch без выражения
    hour := time.Now().Hour()

    fmt.Println("\nПример 2: Время дня")
    switch {
    case hour < 12:
        fmt.Println("Доброе утро!")
    case hour < 18:
        fmt.Println("Добрый день!")
    default:
        fmt.Println("Добрый вечер!")
    }

    // Пример 3: Использование fallthrough
    fmt.Println("\nПример 3: Использование fallthrough")
    num := 2
    switch num {
    case 1:
        fmt.Println("Один")
    case 2:
        fmt.Println("Два")
        fallthrough // Переход к следующему case даже если условие выполнено
    case 3:
        fmt.Println("Три")
    default:
        fmt.Println("Число не распознано")
    }

    // Пример 4: Использование интерфейсов и типов в switch
    fmt.Println("\nПример 4: Switch по типам данных (Type Switch)")
    var i interface{} = "Привет"
    switch v := i.(type) {
    case int:
        fmt.Printf("Это целое число: %d\n", v)
    case string:
        fmt.Printf("Это строка: %s\n", v)
    default:
        fmt.Println("Неизвестный тип")
    }
}
