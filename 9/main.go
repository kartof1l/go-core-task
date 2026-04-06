package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	unsigned := make(chan uint8)   // небуферизированный
	floated := make(chan float64)  // небуферизированный

	wg.Add(1)
	go func() {// ГОРУТИНА КОТОРАЯ ПИШЕТ В КАНАЛ
		defer wg.Done()
		defer close(unsigned)
		for i := uint8(0); i <= 100; i++ {
			unsigned <- i // отправляем в канал
		}
	}()

	wg.Add(1)
	go func() {// ГОРУТИНА КОТОРАЯ ЧИТАЕТ И ПИШЕТ В ФЛОАТ
		defer wg.Done()
		defer close(floated)
		for i := range unsigned { // читаем из unsigned
			result := math.Pow(float64(i), 3)
			fmt.Printf("uint ЗАПИСАН(%d) = %.0f\n", i, result)
			floated <- result 
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for f := range floated {
			fmt.Printf("получен флоат: %.0f\n", f)// читаю из канала сразу т.к. без буфера флоат канал
		}
	}()

	wg.Wait()
}