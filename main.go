package main

import (
    "fmt"
    "sync"
)

func main() {
    // Проблемный вариант (без гонки, но с логической ошибкой)
    fmt.Println("=== Проблемный вариант (все горутины выведут 5) ===")
    var wg1 sync.WaitGroup
    wg1.Add(5)
    
    for i := 0; i < 5; i++ {
        go func() {
            fmt.Println(i)  // захват переменной цикл
            wg1.Done()
        }()
    }
    wg1.Wait()
    
    // Исправленный вариант (передаём копию)
    fmt.Println("\n=== Исправленный вариант (правильные значеня) ===")
    var wg2 sync.WaitGroup
    wg2.Add(5)
    
    for i := 0; i < 5; i++ {
        go func(val int) {  // передаём копию
            fmt.Println(val)
            wg2.Done()
        }(i)
    }
    wg2.Wait()
    
    // Демонстрация реальной гонки данных
    fmt.Println("\n=== Демонстрация гонки данных (закомментировано) ===")
    // Раскомментируйте, чтобы увидеть гонку:
    // demonstrateRace()
}

func MaxInt(a, b int) int {
    if a >= b {
        return a
    }
    return b
}

// Функция для демонстрации реальной гонки данных
func demonstrateRace() {
    var wg sync.WaitGroup
    counter := 0
    
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            counter++  // <- ЭТО ВЫЗЫВАЕТ ГОНКУ ДАННЫХ!
            wg.Done()
        }()
    }
    
    wg.Wait()
    fmt.Println("Counter:", counter)  // результат будет непредсказуемым
}