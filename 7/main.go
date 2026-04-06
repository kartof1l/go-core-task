package main

import (
	"sync"
)

//на вход идет тип раскрытого массива каналов, на выходи один канал
func Merge(channels ...<-chan int) <-chan int {
    out := make(chan int)
    
    var wg sync.WaitGroup
    
    for _, ch := range channels {
        wg.Add(1)
        go func(c <-chan int) {//горутины которые пишут
            defer wg.Done()
            for val := range c {
                out <- val
            }
        }(ch)
    }
    
    go func() {
        wg.Wait()
        close(out)//горутина ждет пока цикл закончится и закрывает, в тестах канал не блокируется т.к. аут будет постоянно читаться
    }()
    
    return out
}