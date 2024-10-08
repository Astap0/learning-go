package main

import "fmt"

func main() {
    x := 10         // Обычная переменная
    p := &x        // & — получаем адрес x

    fmt.Println("x:", x)     // 10
    fmt.Println("p:", p)     // адрес переменной x
    fmt.Println("*p:", *p)   // 10 (разыменование указателя p)

    *p = 20        // Изменяем значение x через указатель
    fmt.Println("x после изменения:", x) // 20
}

// Как запомнить
// Вот несколько советов, чтобы не путаться:

// & для адреса: Запомните, что & означает "где это находится?" (где находится переменная в памяти). Оно помогает получить адрес.
// * для значения: Запомните, что * означает "что там внутри?" (какое значение по этому адресу?). Оно помогает получить значение из указателя.