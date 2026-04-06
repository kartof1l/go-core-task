package main

import (
	"testing"
	"time"
)

func TestWaitGroup(t *testing.T) {
    wg := NewWaitGroup()
    
    var result int
    
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(n int) {
            defer wg.Done()
            result += n
        }(i)
    }
    
    wg.Wait()
    
    if result != 45 {
        t.Errorf("expected 45, got %d", result)
    }
}

func TestWaitGroupWithAddInside(t *testing.T) {
    wg := NewWaitGroup()
    
    var counter int
    
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            wg.Add(1)
            go func() {
                defer wg.Done()
                counter++
            }()
        }()
    }
    
    wg.Wait()
    
    if counter != 5 {
        t.Errorf("expected 5, got %d", counter)
    }
}

func TestWaitGroupParallel(t *testing.T) {
    wg := NewWaitGroup()
    
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            time.Sleep(time.Millisecond)
        }()
    }
    
    wg.Wait()
}

func TestWaitGroupZeroAdd(t *testing.T) {
    wg := NewWaitGroup()
    wg.Wait()
}