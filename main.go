package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    counter := 0  // общая переменная
    
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            counter++  // <- РЕАЛЬНАЯ ГОНКА! несколько горутин пишут одновременн
            wg.Done()
        }()
    }
    
    wg.Wait()
    fmt.Println("Counter:", counter)
}

func MaxInt(a, b int) int {
    if a >= b {
        return a
    }
    return b
}