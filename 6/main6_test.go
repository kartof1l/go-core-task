package main

import (
	"testing"
	"time"
)

func TestRandomGenerator(t *testing.T) {
    ch := RandomGenerator()
    
    for i := 0; i < 10; i++ {
        select {
        case val := <-ch:
            if val != i {
                t.Errorf("expected %d, got %d", i, val)
            }
        case <-time.After(time.Second):
            t.Fatal("timeout")
        }
    }
}

func TestGeneratorCloses(t *testing.T) {
    ch := RandomGenerator()
    
    for i := 0; i < 5; i++ {
        <-ch
    }
    
    for i := 0; i < 10; i++ {
        _, ok := <-ch
        if !ok {
            t.Error("channel closed unexpectedly")
        }
    }
}