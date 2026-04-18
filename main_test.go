package main

import (
    "sync"
    "testing"
)

// ЭТОТ ТЕСТ ПОКАЖЕТ ГОНКУ!
func TestWithRealRace(t *testing.T) {
    var wg sync.WaitGroup
    counter := 0
    
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            counter++ // ← ПРОБЛЕМА: конкурентная запись
            wg.Done()
        }()
    }
    
    wg.Wait()
}