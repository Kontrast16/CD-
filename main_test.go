package main

import (
    "sync"
    "testing"
)

// ЭТОТ ТЕСТ ПОКАЖЕТ ГОНКУ!

// А ЭТОТ ТЕСТ НЕ ПОКАЖЕТ ГОНКУ (как ваш код)
func TestWithoutRace(t *testing.T) {
    var wg sync.WaitGroup
    
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func() {
            _ = i // только чтение, нет записи
            wg.Done()
        }()
    }
    wg.Wait()
}