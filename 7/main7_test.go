package main

import (
	"testing"
)

func TestMerge(t *testing.T) {
    ch1 := make(chan int)
    ch2 := make(chan int)
    ch3 := make(chan int)
    
    go func() {
        ch1 <- 1
        ch1 <- 2
        close(ch1)
    }()
    
    go func() {
        ch2 <- 3
        ch2 <- 4
        close(ch2)
    }()
    
    go func() {
        ch3 <- 5
        ch3 <- 6
        close(ch3)
    }()
    
    merged := Merge(ch1, ch2, ch3)
    
    result := []int{}
    for val := range merged {//тут они постоянно читаются и аппендятся в резалт
        result = append(result, val)
    }
    
    expected := []int{1,2,3,4,5,6}
    if len(result) != len(expected) {
        t.Errorf("got %d elements, want %d", len(result), len(expected))
    }
}

func TestMergeEmpty(t *testing.T) {
    ch := make(chan int)
    close(ch)
    
    merged := Merge(ch)
    
    count := 0
    for range merged {
        count++
    }
    
    if count != 0 {
        t.Errorf("expected 0, got %d", count)
    }
}

func TestMergeSingle(t *testing.T) {
    ch := make(chan int)
    
    go func() {
        ch <- 42
        close(ch)
    }()
    
    merged := Merge(ch)
    
    val := <-merged
    if val != 42 {
        t.Errorf("expected 42, got %d", val)
    }
}